import React, { useState } from "react";
import "./Login.css";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [hover, setHover] = useState(false);
  const [maskedPassword, setMaskedPassword] = useState("");

  const handlePasswordChange = (e) => {
    const value = e.target.value;
    setPassword(value);
    setMaskedPassword("*".repeat(value.length));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    alert(`Имя пользователя: ${username}\nПароль: ${password}`);
  };

  return (
    <div className="terminal-container">
      {/* ASCII-Заголовок */}
      <pre className="ascii-art">
        {`
 ██████╗  ██████╗  ██████╗ ████████╗ ██████╗ ██╗    ██╗
██╔════██╗██╔══██╗██╔═══██╗╚══██╔══╝██╔═══██╗██║    ██║
██║    ██║██████╔╝██║   ██║   ██║   ██║   ██║█████████║
██║    ██║██╔═══╝ ██║   ██║   ██║   ██║   ██║██╔════██║
██║    ██║██║     ╚██████╔╝   ██║   ╚██████╔╝██║    ██║
╚═╝    ╚═╝╚═╝      ╚═════╝    ╚═╝    ╚═════╝ ╚═╝    ╚═╝
        `}
      </pre>

      {/* ASCII-Рамка и Форма */}
      <div className="ascii-border">
        <form onSubmit={handleSubmit} className="login-form">
          <label>
            ┃ ЛОГИН: <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              placeholder="Введите логин"
            /> ┃
          </label>

          <label>
            ┃ ПАРОЛЬ: <input
              type="text"
              value={maskedPassword}
              onChange={handlePasswordChange}
              placeholder="Введите пароль"
            /> ┃
          </label>

          <button
            type="submit"
            className="submit-button"
            onMouseEnter={() => setHover(true)}
            onMouseLeave={() => setHover(false)}
          >
            {hover ? "[ ВВОД ]" : "ВВОД"}
          </button>
        </form>
      </div>
    </div>
  );
};

export default Login;
