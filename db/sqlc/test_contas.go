package db

import (
	"context"
	"testing"

	"Projetos/simplebank/util"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateContaParams{
		Proprietario: util.RandomOwner(),
		Saldo:        util.RandomMoney(),
		Moeda:        util.RandomCurrency(),
	}

	contas, err := testQueries.CreateConta(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, contas)

	require.Equal(t, arg.Proprietario, contas.Proprietario)
	require.Equal(t, arg.Saldo, contas.Saldo)
	require.Equal(t, arg.Moeda, contas.Moeda)

	require.NotZero(t, contas.ID)
	require.NotZero(t, contas.CreatedAt)
}
