const express = require('express');
const app = express();
const PORT = process.env.PORT || 8080;

// Middleware para permitir CORS
app.use((req, res, next) => {
  res.setHeader('Access-Control-Allow-Origin', '*');
  res.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS');
  res.setHeader('Access-Control-Allow-Headers', 'Content-Type');
  if (req.method === 'OPTIONS') {
    res.sendStatus(200);
    return;
  }
  next();
});

// Inicialización de la lista de usuarios
let users = [];

// Ruta para obtener todos los usuarios
app.get('/users', (req, res) => {
  res.json({ users });
});

// Ruta para crear un nuevo usuario
app.post('/users', (req, res) => {
  const newUser = req.body;
  if (!newUser) {
    return res.status(400).json({ error: 'No se proporcionó ningún usuario' });
  }

  // Asignar un nuevo ID al usuario
  newUser.id = users.length + 1;

  // Agregar el nuevo usuario a la lista
  users.push(newUser);

  // Devolver el nuevo usuario como respuesta
  res.status(201).json(newUser);
});

// Iniciar el servidor
app.listen(PORT, () => {
  console.log(`Servidor en ejecución en el puerto ${PORT}`);
});
