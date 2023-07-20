package sync

import (
	"encoding/binary"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis/v8"
)

type ObjSync struct {
	mu           sync.Mutex
	redisClient  *redis.Client
	prvTimestamp int64
	instanceId   int
}

func NewObjSync(redis *redis.Client, instanceId int) *ObjSync {
	obj := ObjSync{
		mu:           sync.Mutex{},
		redisClient:  redis,
		prvTimestamp: 0,
		instanceId:   instanceId,
	}
	return &obj
}
func (oSync *ObjSync) GenServiceObjID() int64 {
	var ret int64 = 0

	binsID := make([]byte, 8)
	baseB := make([]byte, 8)
	instanceB := make([]byte, 4)

	var instanceMod = oSync.instanceId % 256 // max 256 instances

	t := time.Now().UnixMilli()
	prvTimestamp := atomic.LoadInt64(&oSync.prvTimestamp)
	for {
		if t <= prvTimestamp {
			t = prvTimestamp + 1
		}
		if atomic.CompareAndSwapInt64(&oSync.prvTimestamp, prvTimestamp, t) {
			break
		}
		prvTimestamp = atomic.LoadInt64(&oSync.prvTimestamp)
	}

	binary.BigEndian.PutUint64(baseB, uint64(t))
	binary.BigEndian.PutUint32(instanceB, uint32(instanceMod))

	// set first 6 bytes
	binsID[1] = baseB[2]
	binsID[2] = baseB[3]
	binsID[3] = baseB[4]
	binsID[4] = baseB[5]
	binsID[5] = baseB[6]
	binsID[6] = baseB[7]

	// next 1 byte for instance id
	binsID[7] = instanceB[3]

	ret = int64(binary.BigEndian.Uint64(binsID))

	return ret
}
