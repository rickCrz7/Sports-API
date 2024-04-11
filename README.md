<h1>Simple Sports Score Tracker API (Go & MySQL)</h1>
This lightweight API and admin panel combo helps you track sports game scores. Perfect for a mobile app or personal project!

<h3>Features:</h3>

Track game scores for your favorite sports.
Admin panel for easy data management (add, update, delete games) created with Vue.

<h3>Tech Stack:</h3>

Built with Go and MySQL for a clean and efficient backend.

<h2>Getting Started</h2>
<p>Clone repository</p>
<h2>Set Up Database</h2>
<h3>Prerequisites</h3>
<ul>
<li>MySQL server installed and running</li>
<li>MySQL client tool or access to the MySQL command line</li>
<li>MySQL username and password</li>
</ul>

<h4>Steps</h4>
<h4>1. Open 4 MySQL client:</h4>
<p

Command Line: Open a terminal or command prompt and type ```mysql -u your_username -p``` (replace your_username with your actual username).<br>
Graphical Client: Open your preferred MySQL client tool (e.g., MySQL Workbench). </p>
<h4>2. Connect to your MySQL server:</h4>
<p>
Enter your password when prompted.
</p>

<h4>3. Run the SQL script:</h4>
<p>

Command Line: Type the following command and press Enter: <br>

Bash
```
source path/to/data.sql
```
<br> Replace ```path/to/data.sql``` with the actual path to the ```sports_mysql.sql``` file. <br>
Graphical Client: Use the client's "Open file" or "Execute script" option to select and run the ```sports_mysql.sql``` file. </p>
<h4>4. Verify the database and tables:</h4>
<p>

Type ```SHOW DATABASES;``` to check if the ```sports``` database was created successfully.<br>
Type ```SHOW TABLES FROM sports;``` to view the tables created within the ```sports``` database. </p>
<h3>Additional Notes</h3>
<ul>
<li>Ensure you have the necessary permissions to create databases and tables if you encounter errors.</li>
<li>Edit the sports_mysql.sql file before running it again to modify the database structure.</li>
<li>Refer to the official MySQL documentation https://www.mysql.com/ for more advanced MySQL operations.</li>
</ul>

<h2>Run backend</h2>
<h3>Prerequisites</h3>
  <ul>
    <li>Git installed and configured</li>
    <li>Go installed (Download from: <a href="https://go.dev/doc/install">https://go.dev/doc/install</a>)</li>
    <li>Cloned Sports-API repository</li>
  </ul>

  <h4>Steps</h4>
  <h4>1. Navigate to the project directory:</h4>
  <p>
    Use `cd` command to change directories to the cloned repository location.
  </p>

  <h4>2. Install dependencies:</h4>
  <p>Run `go mod download` to download required packages.</p>

  <h4>3. Build the project (optional):</h4>
  <p>Run `go build` to build the backend (refer to project documentation for specific commands).</p>

  <h4>4. Run the project:</h4>
    <p>Run `go run main.go` to run backend. </p>





