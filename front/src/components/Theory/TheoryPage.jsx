// TheoryPage.jsx
import React, { useState } from 'react';
import SidebarMenu from '../Common/SidebarMenu';
import './TheoryPage.css';

const TheoryPage = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [selectedCourse, setSelectedCourse] = useState(null);
  const [selectedLesson, setSelectedLesson] = useState(null);
  const [completedLessons, setCompletedLessons] = useState([]);

  const courses = [
    { id: 1, name: 'Основы Python' },
    { id: 2, name: 'Веб-разработка' },
    { id: 3, name: 'Алгоритмы и структуры данных' }
  ];

  const lessons = {
    1: [
      {
        id: 101,
        title: 'Введение в Python',
        description: 'Основные концепции и синтаксис языка Python',
        content: `
          <h2>Введение в Python</h2>
          <p>Python — это высокоуровневый язык программирования общего назначения, который широко используется в веб-разработке, анализе данных, искусственном интеллекте и автоматизации.</p>
          
          <h3>Основные особенности Python:</h3>
          <ul>
            <li>Простой и понятный синтаксис</li>
            <li>Динамическая типизация</li>
            <li>Интерпретируемый язык</li>
            <li>Большое количество библиотек</li>
            <li>Кроссплатформенность</li>
          </ul>
          
          <h3>Первый код на Python</h3>
          <p>Традиционно первая программа на любом языке — это "Hello, World!":</p>
          <pre><code>print("Hello, World!")</code></pre>
          
          <h3>Переменные и типы данных</h3>
          <p>В Python не нужно объявлять тип переменной:</p>
          <pre><code>name = "Alice"  # Строка
age = 25      # Целое число
height = 1.75 # Число с плавающей точкой
is_student = True # Логическое значение</code></pre>
        `,
        duration: '15 мин'
      },
      {
        id: 102,
        title: 'Условные операторы',
        description: 'Использование if, elif, else для управления потоком выполнения',
        content: `
          <h2>Условные операторы в Python</h2>
          <p>Условные операторы позволяют выполнять разные блоки кода в зависимости от условий.</p>
          
          <h3>Базовый синтаксис if</h3>
          <pre><code>if условие:
    # выполнить этот блок, если условие истинно
    print("Условие истинно")</code></pre>
          
          <h3>Полная форма if-elif-else</h3>
          <pre><code>if x > 10:
    print("x больше 10")
elif x > 5:
    print("x больше 5, но меньше или равно 10")
else:
    print("x меньше или равно 5")</code></pre>
          
          <h3>Логические операторы</h3>
          <p>Для комбинирования условий можно использовать:</p>
          <ul>
            <li><code>and</code> — логическое И</li>
            <li><code>or</code> — логическое ИЛИ</li>
            <li><code>not</code> — логическое НЕ</li>
          </ul>
        `,
        duration: '20 мин',
        locked: true
      }
    ],
    2: [
      {
        id: 201,
        title: 'Основы HTML и CSS',
        description: 'Создание структуры веб-страниц и их стилизация',
        content: `
          <h2>Основы HTML</h2>
          <p>HTML (HyperText Markup Language) — это стандартный язык разметки для создания веб-страниц.</p>
          
          <h3>Базовая структура HTML-документа</h3>
          <pre><code>&lt;!DOCTYPE html&gt;
&lt;html&gt;
&lt;head&gt;
    &lt;title&gt;Моя страница&lt;/title&gt;
&lt;/head&gt;
&lt;body&gt;
    &lt;h1&gt;Заголовок&lt;/h1&gt;
    &lt;p&gt;Абзац текста&lt;/p&gt;
&lt;/body&gt;
&lt;/html&gt;</code></pre>
          
          <h2>Основы CSS</h2>
          <p>CSS (Cascading Style Sheets) — язык стилей, определяющий отображение HTML-документов.</p>
          
          <h3>Пример CSS-правила</h3>
          <pre><code>h1 {
    color: blue;
    font-size: 24px;
    text-align: center;
}</code></pre>
        `,
        duration: '25 мин'
      }
    ],
    3: [
      {
        id: 301,
        title: 'Сложность алгоритмов',
        description: 'Оценка времени выполнения алгоритмов',
        content: `
          <h2>Сложность алгоритмов</h2>
          <p>Сложность алгоритма — это оценка количества ресурсов (времени и памяти), необходимых для его выполнения.</p>
          
          <h3>O-нотация (нотация "О-большое")</h3>
          <p>Используется для описания асимптотического поведения функций. В анализе алгоритмов описывает скорость роста времени выполнения при увеличении размера входных данных.</p>
          
          <h3>Основные классы сложности</h3>
          <ul>
            <li><strong>O(1)</strong> — константное время (доступ к элементу массива)</li>
            <li><strong>O(log n)</strong> — логарифмическое время (бинарный поиск)</li>
            <li><strong>O(n)</strong> — линейное время (поиск в неотсортированном массиве)</li>
            <li><strong>O(n log n)</strong> — быстрая сортировка, сортировка слиянием</li>
            <li><strong>O(n²)</strong> — квадратичное время (пузырьковая сортировка)</li>
            <li><strong>O(2ⁿ)</strong> — экспоненциальное время (решение задачи коммивояжера полным перебором)</li>
          </ul>
        `,
        duration: '30 мин'
      }
    ]
  };

  const handleStartLesson = (lesson) => {
    setSelectedLesson(lesson);
  };

  const handleCompleteLesson = () => {
    if (selectedLesson && !completedLessons.includes(selectedLesson.id)) {
      setCompletedLessons([...completedLessons, selectedLesson.id]);
    }
    setSelectedLesson(null);
  };

  return (
    <div className="theory-layout">
      <button 
        className="menu-toggle-button" 
        onClick={() => setIsMenuOpen(true)}
      >
        {'☰'}
      </button>

      <SidebarMenu 
        isOpen={isMenuOpen} 
        onClose={() => setIsMenuOpen(false)} 
      />

      <div className="theory-main-content">
        <header className="theory-header">
          <h1>Теоретические материалы</h1>
          <p>Изучайте теорию и закрепляйте знания на практике</p>
        </header>

        {!selectedLesson ? (
          <>
            <div className="course-selector-container">
              <select
                className="course-selector"
                value={selectedCourse || ''}
                onChange={(e) => setSelectedCourse(e.target.value ? parseInt(e.target.value) : null)}
              >
                <option value="">Выберите курс</option>
                {courses.map(course => (
                  <option key={course.id} value={course.id}>{course.name}</option>
                ))}
              </select>
            </div>

            {selectedCourse && (
              <div className="lessons-container">
                <h2 className="lessons-title">Материалы по курсу: {courses.find(c => c.id === selectedCourse).name}</h2>
                <div className="lessons-grid">
                  {lessons[selectedCourse]?.map(lesson => (
                    <div 
                      key={lesson.id} 
                      className='lesson-card'
                    >
                      <div className="lesson-content">
                        <div className="lesson-header">
                          <h3 className="lesson-title">{lesson.title}</h3>
                          {completedLessons.includes(lesson.id) && (
                            <span className="lesson-completed-badge">✓</span>
                          )}
                        </div>
                        <p className="lesson-description">{lesson.description}</p>
                        <div className="lesson-meta">
                          <span className="lesson-duration">{lesson.duration}</span>
                        </div>
                      </div>
                      <div className="lesson-actions">
                        <button
                          className='start-lesson-button'
                          onClick={() => handleStartLesson(lesson)}
                        >
                          {completedLessons.includes(lesson.id) ? 'Повторить урок' : 'Начать урок'}
                        </button>
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            )}
          </>
        ) : (
          <div className="lesson-content-container">
            <button 
              className="back-button"
              onClick={handleCompleteLesson}
            >
              ← Назад к списку уроков
            </button>

            <div className="lesson-details">
              <h2 className="lesson-title">{selectedLesson.title}</h2>
              <p className="lesson-description">{selectedLesson.description}</p>
              
              <div 
                className="theory-content"
                dangerouslySetInnerHTML={{ __html: selectedLesson.content }}
              />
              
              <div className="lesson-complete-section">
                <button 
                  className="complete-lesson-button"
                  onClick={handleCompleteLesson}
                >
                  {completedLessons.includes(selectedLesson.id) ? 'Урок завершен' : 'Завершить урок'}
                </button>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default TheoryPage;