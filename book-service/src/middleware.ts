import { Request, Response, NextFunction } from 'express';
import jwt from 'jsonwebtoken';

const secretKey = process.env.SECRET_KEY || '';
export interface DecodedToken {
  username: string; // Adjust this according to your JWT payload structure
  role: string;
  exp: number;
}
export interface CustomRequest extends Request {
  decodedToken: DecodedToken;
}
export function verifyToken(req: Request, res: Response, next: NextFunction) {
  const token = req.headers.authorization?.split(' ')[1];

  if (!token) {
    return res.status(401).json({ error: 'Unauthorized: Missing token' });
  }

  try {
    const decoded = jwt.verify(token, secretKey, {
      complete: false,
    }) as DecodedToken;
    (req as CustomRequest).decodedToken = decoded;

    next();
  } catch (error) {
    return res.status(401).json({ error: 'Unauthorized: Invalid token' });
  }
}

export function checkRole(req: Request, res: Response, next: NextFunction) {
  const { decodedToken } = (req as CustomRequest);
  if (decodedToken.role !== 'LIBRARIAN') {
    return res.status(403).json({ error: 'Forbidden: Insufficient privileges' });
  }


  next();
}