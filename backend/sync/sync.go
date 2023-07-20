package sync

import (
	"encoding/binary"
	"sync"
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

	var instanceMod = oSync.instanceId % 256 // max 256 instance

	oSync.mu.Lock()
	defer oSync.mu.Unlock()

	t := time.Now().UnixMilli()
	if t <= oSync.prvTimestamp {
		ret = oSync.prvTimestamp + 1
	} else {
		ret = t
	}
	oSync.prvTimestamp = ret

	binary.BigEndian.PutUint64(baseB, uint64(ret))
	binary.BigEndian.PutUint32(instanceB, uint32(instanceMod))

	// set first 6byte
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
