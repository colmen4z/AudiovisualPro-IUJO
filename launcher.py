import subprocess
import os

npm_path = r"C:\Program Files\nodejs\npm.cmd"

backend_dir = os.path.join("backend")
frontend_dir = os.path.join("frontend")

backend_process = subprocess.Popen(["go", "run", "main.go"], cwd=backend_dir)
print("Backend (Go) iniciado...")

frontend_process = subprocess.Popen([npm_path, "run", "dev"], cwd=frontend_dir)
print("Frontend (Vue + Vite) iniciado...")

backend_process.wait()
frontend_process.wait()
