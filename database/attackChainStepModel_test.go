package database_test

// func TestGetAttackChainSteps(t *testing.T) {
// 	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
// 	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer pgPool.Close()
// 	dbManager := &database.DBManager{DBPool: pgPool}
// 	attackChainSteps, _ := dbManager.GetAttackChainSteps(businessId)

// 	for _, attackChainStep := range attackChainSteps {
// 		assert.IsEqual(attackChainStep.BusinessID.String(), businessId)
// 	}
// }

// func TestGetAttackChainStep(t *testing.T) {
// 	var actionId = "535705bc-fddb-4e2a-8c1c-196755ce16b6"
// 	var attackChainId = "535705bc-fddb-4e2a-8c1c-196755ce16b6"
// 	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
// 	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer pgPool.Close()
// 	dbManager := &database.DBManager{DBPool: pgPool}
// 	attackChainStep, _ := dbManager.GetAttackChainStep(actionId, attackChainId)

// 	assert.IsEqual(attackChainStep.ActionID.String(), actionId)
// 	assert.IsEqual(attackChainStep.AttackChainID.String(), attackChainId)
// }

// func TestDeleteAttackChainStep(t *testing.T) {
// 	var actionId = "535705bc-fddb-4e2a-8c1c-196755ce16b6"
// 	var attackChainId = "535705bc-fddb-4e2a-8c1c-196755ce16b6"
// 	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
// 	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer pgPool.Close()
// 	dbManager := &database.DBManager{DBPool: pgPool}
// 	attackChainStepInput := database.AttackChainStepModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
// 	createAttackChainStepOutput, _ := dbManager.CreateAttackChainStep(attackChainStepInput)

// 	tempAttackChainStepId := createAttackChainStepOutput.ID.String()

// 	err = dbManager.DeleteAttackChainStep(tempAttackChainStepId)

// 	assert.Equal(t, err, nil)

// 	_, err = dbManager.GetAttackChainStep(tempAttackChainStepId)

// 	assert.NotEqual(t, err, nil)

// }

// func TestCreateAttackChainStep(t *testing.T) {
// 	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
// 	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer pgPool.Close()
// 	dbManager := &database.DBManager{DBPool: pgPool}
// 	attackChainStepInput := database.AttackChainStepModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
// 	createAttackChainStepOutput, err := dbManager.CreateAttackChainStep(attackChainStepInput)

// 	assert.Equal(t, err, nil)

// 	attackChainStep, err := dbManager.GetAttackChainStep(createAttackChainStepOutput.ID.String())

// 	assert.Equal(t, err, nil)

// 	assert.Equal(t, attackChainStep, createAttackChainStepOutput)
// }

// func TestUpdateAttackChainStep(t *testing.T) {

// 	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
// 	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer pgPool.Close()
// 	dbManager := &database.DBManager{DBPool: pgPool}
// 	createAttackChainStepInput := database.AttackChainStepModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
// 	createAttackChainStepOutput, _ := dbManager.CreateAttackChainStep(createAttackChainStepInput)

// 	createAttackChainStepOutput.Name = "test2"

// 	updateAttackChainStepInput := createAttackChainStepOutput

// 	err = dbManager.UpdateAttackChainStep(updateAttackChainStepInput)

// 	assert.Equal(t, err, nil)

// 	updatedAttackChainStep, _ := dbManager.GetAttackChainStep(createAttackChainStepOutput.ID.String())

// 	assert.Equal(t, updateAttackChainStepInput.Name, updatedAttackChainStep.Name)
// }
