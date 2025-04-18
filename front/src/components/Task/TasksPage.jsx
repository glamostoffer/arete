// TasksPage.jsx
import React, { useState } from 'react';
import SidebarMenu from '../Common/SidebarMenu';
import './TasksPage.css';

const TasksPage = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [selectedCourse, setSelectedCourse] = useState(null);
  const [selectedTask, setSelectedTask] = useState(null);
  const [code, setCode] = useState('');
  const [output, setOutput] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [completedTasks, setCompletedTasks] = useState([]);

  // Курсы с задачами (только программирование)
  const programmingCourses = [
    { id: 1, name: 'Основы Python' },
    { id: 3, name: 'Алгоритмы и структуры данных' },
    { id: 6, name: 'Основы баз данных' }
  ];

  // Задачи по курсам
  const tasks = {
    1: [
      {
        id: 101,
        title: 'Сумма двух чисел',
        description: 'Напишите функцию, которая принимает два числа и возвращает их сумму.',
        difficulty: 'Легкая',
        testCases: [
          { input: '2, 3', output: '5' },
          { input: '-1, 1', output: '0' },
          { input: '0, 0', output: '0' }
        ],
        defaultCode: 'def sum_two_numbers(a, b):\n    # Ваш код здесь\n    pass',
        explanation: 'Эта задача проверяет ваше понимание базовых операций в Python. Функция должна просто вернуть сумму двух аргументов.'
      },
      {
        id: 102,
        title: 'Факториал числа',
        description: 'Напишите функцию для вычисления факториала числа n.',
        difficulty: 'Средняя',
        locked: true,
        testCases: [
          { input: '5', output: '120' },
          { input: '0', output: '1' },
          { input: '1', output: '1' }
        ],
        defaultCode: 'def factorial(n):\n    # Ваш код здесь\n    pass',
        explanation: 'Факториал числа n - это произведение всех положительных целых чисел от 1 до n. Учтите особый случай для 0! = 1.'
      }
    ],
    3: [
      {
        id: 301,
        title: 'Бинарный поиск',
        description: 'Реализуйте алгоритм бинарного поиска в отсортированном массиве.',
        difficulty: 'Средняя',
        testCases: [
          { input: '[1, 2, 3, 4, 5], 3', output: '2' },
          { input: '[1, 3, 5, 7, 9], 2', output: '-1' }
        ],
        defaultCode: 'def binary_search(arr, target):\n    # Ваш код здесь\n    pass',
        explanation: 'Бинарный поиск работает за O(log n) времени. Массив должен быть отсортирован. Возвращает индекс элемента или -1, если элемент не найден.'
      }
    ],
    6: [
      {
        id: 601,
        title: 'SQL запрос: Выборка данных',
        description: 'Напишите SQL запрос, который возвращает всех пользователей старше 18 лет.',
        difficulty: 'Легкая',
        testCases: [
          { input: 'Таблица users с полями id, name, age', output: 'Список пользователей с age > 18' }
        ],
        defaultCode: 'SELECT * FROM users WHERE age > 18;',
        explanation: 'Это базовый SQL запрос с условием WHERE для фильтрации данных.'
      }
    ]
  };

  const handleRunCode = () => {
    setIsLoading(true);
    // Здесь должна быть логика отправки кода на сервер для выполнения
    // Для демонстрации просто симулируем ответ
    setTimeout(() => {
      setOutput('Тесты пройдены успешно! 🎉');
      setIsLoading(false);
    }, 1500);
  };

  const handleSubmitSolution = () => {
    setIsLoading(true);
    // Здесь должна быть логика отправки решения на проверку
    setTimeout(() => {
      setOutput('Решение принято! Задача выполнена.');
      setCompletedTasks([...completedTasks, selectedTask.id]);
      setIsLoading(false);
    }, 2000);
  };

  const isTaskUnlocked = (task) => {
    if (!task.locked) return true;
    const courseTasks = tasks[selectedCourse];
    const taskIndex = courseTasks.findIndex(t => t.id === task.id);
    return taskIndex > 0 && completedTasks.includes(courseTasks[taskIndex - 1].id);
  };

  const handleStartTask = (task) => {
    setSelectedTask(task);
    setCode(task.defaultCode);
    setOutput('');
  };

  return (
    <div className="tasks-layout">
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

      <div className="tasks-main-content">
        <header className="tasks-header">
          <h1>Задачи по программированию</h1>
          <p>Решайте задачи, чтобы закрепить знания и подготовиться к собеседованиям</p>
        </header>

        {!selectedTask ? (
          <>
            <div className="course-selector-container">
              <select
                className="course-selector"
                value={selectedCourse || ''}
                onChange={(e) => setSelectedCourse(e.target.value ? parseInt(e.target.value) : null)}
              >
                <option value="">Выберите курс</option>
                {programmingCourses.map(course => (
                  <option key={course.id} value={course.id}>{course.name}</option>
                ))}
              </select>
            </div>

            {selectedCourse && (
              <div className="tasks-container">
                <h2 className="tasks-title">Задачи по курсу: {programmingCourses.find(c => c.id === selectedCourse).name}</h2>
                <div className="tasks-grid">
                  {tasks[selectedCourse]?.map(task => (
                    <div 
                      key={task.id} 
                      className={`task-card ${!isTaskUnlocked(task) ? 'locked' : ''}`}
                    >
                      <div className="task-content">
                        <div className="task-header">
                          <h3 className="task-title">{task.title}</h3>
                          <span className={`task-difficulty ${task.difficulty.toLowerCase()}`}>
                            {task.difficulty}
                          </span>
                        </div>
                        <p className="task-description">{task.description}</p>
                        {!isTaskUnlocked(task) && (
                          <div className="task-lock-message">
                            Пройдите предыдущую задачу для разблокировки
                          </div>
                        )}
                      </div>
                      <div className="task-actions">
                        <button
                          className={`start-task-button ${!isTaskUnlocked(task) ? 'disabled' : ''}`}
                          onClick={() => isTaskUnlocked(task) && handleStartTask(task)}
                          disabled={!isTaskUnlocked(task)}
                        >
                          Решить задачу
                        </button>
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            )}
          </>
        ) : (
          <div className="task-solution-container">
            <button 
              className="back-button"
              onClick={() => setSelectedTask(null)}
            >
              ← Назад к списку задач
            </button>

            <div className="task-details">
              <h2 className="task-title">{selectedTask.title}</h2>
              <span className={`task-difficulty ${selectedTask.difficulty.toLowerCase()}`}>
                {selectedTask.difficulty}
              </span>
              
              <div className="task-description">
                <p>{selectedTask.description}</p>
              </div>

              <div className="test-cases">
                <h3>Примеры тестов:</h3>
                {selectedTask.testCases.map((testCase, index) => (
                  <div key={index} className="test-case">
                    <p><strong>Входные данные:</strong> {testCase.input}</p>
                    <p><strong>Ожидаемый вывод:</strong> {testCase.output}</p>
                  </div>
                ))}
              </div>

              <div className="explanation">
                <h3>Объяснение:</h3>
                <p>{selectedTask.explanation}</p>
              </div>
            </div>

            <div className="code-editor-container">
              <h3>Ваше решение:</h3>
              <textarea
                className="code-editor"
                value={code}
                onChange={(e) => setCode(e.target.value)}
                spellCheck="false"
              />
              <div className="editor-actions">
                <button
                  className="run-button"
                  onClick={handleRunCode}
                  disabled={isLoading}
                >
                  {isLoading ? 'Выполнение...' : 'Запустить код'}
                </button>
                <button
                  className="submit-button"
                  onClick={handleSubmitSolution}
                  disabled={isLoading}
                >
                  {isLoading ? 'Проверка...' : 'Отправить решение'}
                </button>
              </div>
            </div>

            {output && (
              <div className="output-container">
                <h3>Результат:</h3>
                <pre className="output">{output}</pre>
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  );
};

export default TasksPage;