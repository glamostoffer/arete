import React, { useState, useRef } from 'react'
import { useNavigate } from 'react-router-dom'
import './LoginPage.css'
import ky from 'ky'
import { useAuthStore, useSignUpStore } from '../../model/store'
import { useShallow } from 'zustand/react/shallow'

const EmailConfirmationPage = () => {
  const [codeDigits, setCodeDigits] = useState(['', '', '', '', '', ''])
  const [error, setError] = useState('')
  const inputsRef = useRef([])
  const navigate = useNavigate()

  const signUpFinalize = useSignUpStore(useShallow((state) => state.signUpFinalize))

  const handleChange = (value, index) => {
    if (!/^[0-9]?$/.test(value)) return // Только цифры и пусто

    const newDigits = [...codeDigits]
    newDigits[index] = value
    setCodeDigits(newDigits)

    if (value && index < 5) {
      inputsRef.current[index + 1].focus()
    }
  }

  const handleKeyDown = (e, index) => {
    if (e.key === 'Backspace' && !codeDigits[index] && index > 0) {
      inputsRef.current[index - 1].focus()
    }
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    const fullCode = codeDigits.join('')
    signUpFinalize(fullCode)
      .then(() => navigate('/profile'))
      .catch(() => setError('Неверный код. Попробуйте снова.'))
  }

  const isCodeComplete = codeDigits.every((digit) => digit !== '')

  return (
    <div className='page-background'>
      <div className='light-orb'></div>
      <div className='login-container'>
        <div className='login-header'>
          <div className='login-logo'>Арете</div>
        </div>

        <div className='login-content'>
          <h1 className='login-title'>Подтвердите вашу почту</h1>
          <p style={{ textAlign: 'center', marginBottom: '20px', color: '#607d8b' }}>
            Введите 6-значный код, отправленный на вашу почту.
          </p>

          <form className='login-form' onSubmit={handleSubmit}>
            <div
              style={{
                display: 'flex',
                justifyContent: 'space-between',
                gap: '10px',
                marginBottom: '20px',
              }}
            >
              {codeDigits.map((digit, index) => (
                <input
                  key={index}
                  type='text'
                  inputMode='numeric'
                  maxLength='1'
                  className='code-input'
                  value={digit}
                  onChange={(e) => handleChange(e.target.value, index)}
                  onKeyDown={(e) => handleKeyDown(e, index)}
                  ref={(el) => (inputsRef.current[index] = el)}
                />
              ))}
            </div>

            {error && (
              <div style={{ color: 'red', marginBottom: '10px', textAlign: 'center' }}>
                {error}
              </div>
            )}

            <button
              type='submit'
              className={`login-button ${!isCodeComplete ? 'login-button-disabled' : ''}`}
              disabled={!isCodeComplete}
            >
              Подтвердить
            </button>
          </form>
        </div>

        <div className='login-footer'>
          <p className='login-footer-text'>
            Не получили код? Проверьте папку "Спам" или попробуйте отправить снова.
          </p>
        </div>
      </div>
    </div>
  )
}

export default EmailConfirmationPage
