# api-students
API to 'Golang do Zero' course students

Routes:
- GET /students - List all students
- POST /students - Create student
- GET /students/:id Get infos from a especific student
- PUT /students/:id - Update student
- DELETE /students/:id - Delete
- GET /students?active<true/false> - List all active/non-active students

Struct Student: 
- Name 
- CPF 
- Email
- Age 
- Active