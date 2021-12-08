package kafka

import (
	"github.com/sSchmidtT/imersaofc5/adapter/presenter/transaction"
	"github.com/sSchmidtT/imersaofc5/domain/entity"
	"github.com/sSchmidtT/imersaofc5/usecase/process_transaction"
	"github.com/stretchr/testify/assert"
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func TestProducerPublish(t *testing.T) {
	expectedOutput := process_transaction.TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you dont have limit for this transaction",
	}
	//outputJson, _ := json.Marshal(expectedOutput)

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}
	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())
	err := producer.Publish(expectedOutput, []byte("1"), "test")
	assert.Nil(t, err)

}
