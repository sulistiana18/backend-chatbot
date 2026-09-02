package ai

import (
	"context"
	"fmt"

	"backend/entity"
)

type dummyProvider struct{}

func NewDummyProvider() entity.AIProvider {
	return &dummyProvider{}
}

// daftar pertanyaan gali informasi, dipakai berurutan sesuai jumlah history
var dummyQuestions = []struct {
	Content string
	Options []string
}{
	{
		Content: "Baik, sebelum mesin mati, apakah lampu dashboard masih menyala saat kunci diputar?",
		Options: []string{"Ya, menyala", "Tidak, mati total", "Tidak yakin"},
	},
	{
		Content: "Apakah terdengar bunyi 'klik-klik' dari mesin saat mencoba starter?",
		Options: []string{"Ya, ada bunyi klik", "Tidak ada bunyi sama sekali"},
	},
	{
		Content: "Sudah berapa lama mobil terakhir digunakan sebelum kejadian ini?",
		Options: nil, // contoh pertanyaan tanpa pilihan, user jawab bebas
	},
}

func (d *dummyProvider) GenerateReply(ctx context.Context, history []entity.Message, newMessage string) (*entity.AIReply, error) {
	assistantCount := 0
	for _, m := range history {
		if m.Role == entity.RoleAssistant {
			assistantCount++
		}
	}

	if assistantCount < len(dummyQuestions) {
		q := dummyQuestions[assistantCount]
		return &entity.AIReply{Content: q.Content, Options: q.Options}, nil
	}

	return &entity.AIReply{
		Content: fmt.Sprintf("Berdasarkan gejala yang kamu sampaikan (%q), kemungkinan besar penyebabnya adalah aki lemah atau starter motor bermasalah.", newMessage),
		Options: nil,
	}, nil
}