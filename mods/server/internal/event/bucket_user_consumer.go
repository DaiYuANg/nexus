package event

import "nexus/internal/entity"

type BucketUserConsumer struct {
}

func (b *BucketUserConsumer) Consumer(user entity.User) {

}
