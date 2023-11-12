import { MongoClient } from 'mongodb';

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

      if (!existingUser) {
        return res.status(401).json({ error: 'Login failed. Please check your credentials.' });
      }

      const passwordMatch = password === existingUser.hashedpassword; // Change this line

      if (!passwordMatch) {
        return res.status(401).json({ error: 'Login failed. Please check your credentials.' });
      }

      const token = 'your_generated_token_here';

      res.status(200).json({ token, user: { email: existingUser.email } });
    } catch (error) {
      console.error('Login error:', error);
      res.status(500).json({ error: 'Login failed. Please try again.' });
    } finally {
      await mongoClient.close();
    }
  } else {
    res.status(405).end(); // Method Not Allowed
  }
}
