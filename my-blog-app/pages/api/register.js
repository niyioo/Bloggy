import { MongoClient } from 'mongodb';
import { hashPassword } from '../../auth/auth';

export default async function handler(req, res) {
  if (req.method === 'POST') {
    const { email, password } = req.body;

    if (!email || !password) {
      return res.status(400).json({ error: 'Email and password are required.' });
    }

    const mongoClient = new MongoClient('mongodb://localhost:27017/');

    try {
      await mongoClient.connect();
      const db = mongoClient.db('bloggy');
      const collection = db.collection('users');

      const existingUser = await collection.findOne({ email });

      if (existingUser) {
        return res.status(400).json({ error: 'Email is already in use. Please use a different email.' });
      }

      const hashedPassword = await hashPassword(password);

      const user = {
        email,
        hashedPassword,
      };

      const result = await collection.insertOne(user);

      if (result.insertedId) {
        res.status(200).json({ message: 'Registration successful.' });
      } else {
        res.status(500).json({ error: 'Registration failed. Please try again.' });
      }
    } catch (error) {
      console.error('Registration error:', error);
      res.status(500).json({ error: 'Registration failed. Please try again.' });
    } finally {
      await mongoClient.close();
    }
  } else {
    res.status(405).end(); // Method Not Allowed
  }
}
