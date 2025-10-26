import { pool } from '../database/database.js'
import bcrypt from 'bcryptjs'
import dotenv from 'dotenv'
import jwt from 'jsonwebtoken'

dotenv.config()
const SECRET = process.env.JWT_SECRET

export const login = async (req, res) => {
    try {
        const { usuario_gestor, pass_gestor } = req.body

        const result = await pool.query(
            'SELECT * FROM gestor WHERE usuario_gestor = $1',
            [usuario_gestor]
        )

        const gestor = result.rows[0]
        if (!gestor) return res.status(404).json({ message: 'Usuario no encontrado.' });

        const valid = await bcrypt.compare(pass_gestor, gestor.pass_gestor)
        if (!valid) return json.status(401).json({ message: 'Contraseña incorrecta.' });

        const token = jwt.sign(
            {
                id_gestor: gestor.id_gestor,
                usuario_gestor: gestor.usuario_gestor,
                nombre_gestor: gestor.nombre_gestor
            },
            SECRET,
            { expiresIn: '2h' }
        )

        res.json({ token })
    } catch (err) {
        res.status(500).json({ message: err.message })
    }
}