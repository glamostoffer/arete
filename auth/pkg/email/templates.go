package email

// message templates
const (
	TemplateConfirmationCode = `
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Подтверждение email</title>
</head>
<body style="font-family: 'Helvetica Neue', Arial, sans-serif; background-color: #000000; margin: 0; padding: 0; color: #ffffff; font-weight: 300; line-height: 1.6;">
    <div class="container" style="max-width: 500px; margin: 20px auto; background-color: #0a0a0a; border-radius: 12px; overflow: hidden; box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3); border: 1px solid #222222;">
        <div class="header" style="padding: 20px; text-align: center; background-color: #000000; border-bottom: 1px solid #222222;">
            <div class="logo" style="font-size: 20px; font-weight: 400; color: #ffffff; letter-spacing: 0.5px;">Арете</div>
        </div>
        
        <div class="content" style="padding: 28px; text-align: center;">
            <h1 style="font-size: 20px; margin-bottom: 20px; color: #ffffff; font-weight: 400;">Подтвердите вашу электронную почту</h1>
            
            <p style="font-size: 15px; line-height: 1.5; margin-bottom: 20px; color: #cccccc; font-weight: 300;">Для завершения регистрации введите следующий код подтверждения:</p>
            
            <div class="code-box" style="background-color: #111111; border-radius: 8px; padding: 12px; margin: 20px auto; display: inline-block; border: 1px solid #333333; box-shadow: 0 1px 4px rgba(0, 0, 0, 0.2); max-width: 80%;">
                <div class="code" style="font-size: 24px; letter-spacing: 6px; font-weight: 300; color: #ffffff; padding: 6px 12px; font-family: 'Courier New', monospace;">%s</div>
            </div>
            
            <p style="font-size: 15px; line-height: 1.5; margin-bottom: 20px; color: #cccccc; font-weight: 300;">Этот код действителен в течение 30 минут. Если вы не запрашивали этот код, просто проигнорируйте это письмо.</p>
            
            <p class="small" style="font-size: 13px; color: #777777; font-weight: 300;">С уважением,<br>Команда Арете</p>
        </div>
        
        <div class="footer" style="padding: 20px; text-align: center; font-size: 13px; color: #888888; background-color: #000000; border-top: 1px solid #222222;">
            <p class="small" style="font-size: 13px; color: #777777; font-weight: 300;">Если у вас возникли вопросы, свяжитесь с нашей поддержкой.</p>
        </div>
    </div>
</body>
</html>
`
)
