package event

import "nexus/internal/entity"

type SendEmailConsumer struct{}

func (b *SendEmailConsumer) Consumer(user entity.User) {

}
