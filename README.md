# AI-Powered Task Management System

A full-stack task management application built with **Golang (Gin)** for the backend and **Next.js (TypeScript, Tailwind, Shadcn UI)** for the frontend. It features user authentication, task CRUD operations, real-time updates via WebSockets, and AI-powered task suggestions using Google’s Gemini API. This project was developed as part of a 4-hour coding challenge, emphasizing speed, functionality, and AI integration.

## Features
- **Authentication:** JWT-based user registration and login.
- **Task Management:** Create, read, and (partially) update tasks with status and assignment options.
- **AI Suggestions:** Context-aware subtask suggestions powered by Gemini AI, based on existing tasks.
- **UI:** modern interface with Shadcn UI components and Tailwind CSS.

## Prerequisites
- **Golang**: 1.21+ (for backend)
- **Node.js**: 18+ (for frontend)
- **PostgreSQL**: NeonDB instance (or local PostgreSQL)
- **Gemini API Key**: For AI suggestions (get from Google Cloud or AI Studio)

## Getting Started

### Backend Setup
1. **Clone the Repository:**
   ```bash
   git clone https://github.com/rudrikprajapati/zocket-task
   ```

2. **Install Dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the Backend:**
   ```bash
   go run main.go
   ```
   - Runs on `http://localhost:8080`.

### Frontend Setup
1. **Navigate to Frontend Directory:**
   ```bash
   cd task-manager-frontend
   ```

2. **Install Dependencies:**
   ```bash
   npm install
   ```

3. **Run the Frontend:**
   ```bash
   npm run dev
   ```
   - Runs on `http://localhost:3000`.

4. **Access the App:**
   - Open `http://localhost:3000` in your browser.
   - Register a user, log in, and start managing tasks.

## My Approach
### Development Strategy
- **Time Management:** Prioritized core features (auth, task CRUD, AI) within 4 hours, allocating ~2 hours for backend and ~2 hours for frontend.
- **Backend:** Used Golang with Gin for its speed and simplicity. Structured with models, controllers, and routes for maintainability. Integrated NeonDB for persistent storage and WebSockets for real-time updates(Leave to setup .env 😅).
- **Frontend:** Chose Next.js with TypeScript for type safety and rapid UI development. Used Shadcn UI and Tailwind CSS for a polished, responsive design.
- **AI Integration:** Leveraged Gemini API to provide context-aware task suggestions, enhancing user productivity.

### How AI Helps (Human-Made Feel)
- **Contextual Suggestions:** The AI analyzes existing tasks (titles, descriptions, statuses) alongside user input to suggest relevant subtasks, mimicking a human assistant’s ability to understand context.
- **Human Oversight:** The UI and workflow (e.g., manual task creation, status selection) ensure users retain control, with AI acting as a supportive tool rather than an automated overlord.
- **Iterative Refinement:** I tested AI outputs locally (e.g., via Postman) used react markdown to render ai generated output
## Current Status
- **Backend:** Fully functional locally (`http://localhost:8080`) with auth, task management, WebSocket, and AI endpoints. Not yet deployed.
- **Frontend:** Runs locally (`http://localhost:3000`) with login, task creation, and AI suggestions.

## Deployment (To-Do)
### Backend (Render)
1. Push to GitHub.
2. Create a `Procfile`: `web: task-manager`.
3. Deploy on Render with env vars (`DATABASE_URL`, `GEMINI_API_KEY`, `JWT_SECRET`).

### Frontend (Vercel)
1. Push to GitHub.
2. Run `vercel --prod`.
3. Update API URLs to point to the deployed backend.

## Future Improvements
- Full CRUD (add `PUT`/`DELETE` for tasks).
- WebSocket authentication fix.
- Input validation in `CreateTaskForm`.
- Deployed demo with video walkthrough.


https://github.com/user-attachments/assets/48952216-ef94-4c8b-8d4f-d4356a10fa07


![90F812A9-E867-47A4-8706-D4AFB7D6A697_1_102_o](https://github.com/user-attachments/assets/085dc6bd-bfeb-4ac3-983c-7169f7c922a7)

