package flowtest

import "fmt"

type AccountFlowContext struct {
	Email   string
	UserID  int
	Success bool
}

func CreateAccount(ctx AccountFlowContext) (AccountFlowContext, error) {
	fmt.Println("✅ Creating account for:", ctx.Email)

	// 假設建立成功並回傳 userID
	ctx.UserID = 12345
	ctx.Success = true
	return ctx, nil
}

func NotifyAccountCreated(ctx AccountFlowContext) (AccountFlowContext, error) {
	if !ctx.Success {
		return ctx, fmt.Errorf("⛔ Cannot notify: account creation failed")
	}
	fmt.Printf("📧 Sending notification for user %d (%s)\n", ctx.UserID, ctx.Email)
	return ctx, nil
}

func CreateJiraTask(ctx AccountFlowContext) (AccountFlowContext, error) {
	if !ctx.Success {
		return ctx, fmt.Errorf("⛔ Cannot create JIRA task: user creation failed")
	}
	fmt.Printf("🛠 Creating JIRA task for user %d\n", ctx.UserID)
	return ctx, nil
}

func RunCreateAccountFlow(email string) error {
	ctx := AccountFlowContext{
		Email: email,
	}

	var err error
	ctx, err = CreateAccount(ctx)
	if err != nil {
		return err
	}

	ctx, err = NotifyAccountCreated(ctx)
	if err != nil {
		return err
	}

	ctx, err = CreateJiraTask(ctx)
	if err != nil {
		return err
	}

	fmt.Println("✅ 全部流程完成")
	return nil
}
