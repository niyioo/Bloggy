import { MongoClient } from 'mongodb';
import bcrypt from 'bcrypt';

async function hashPassword(password) {
  const saltRounds = 10;
  const hashedPassword = await bcrypt.hash(password, saltRounds);
  return hashedPassword;
}

export default async function handler(req, res) {
  if (req.method === 'POST') {
    const { email, password } = req.body;

    if (!email || !password) {
      return res.status(400).json({ error: 'Email and password are required.' });
    }

    const mongoClient = new MongoClient('mongodb://localhost:27017/bloggy');

    try {
      await mongoClient.connect();
      console.log('Connected to MongoDB');
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
