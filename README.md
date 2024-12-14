# HakikaRide

## Overview
**HakikaRide** is a real-time school bus tracking application designed to ensure the safety and efficiency of school transportation. This project allows parents, schools, and drivers to monitor bus locations, ensure timely pickups and drop-offs, and foster peace of mind. Built as an MVP for a hackathon, this system integrates **Go** for the backend, **SQLite** as the database, and features a user-friendly frontend built with **HTML, CSS, and JavaScript**.

## Features
1. **Parent Dashboard**: 
   - Track real-time bus location.
   - View child’s pickup/drop-off status.
   - Notifications for delays or route changes.

2. **Driver Dashboard**:
   - Update bus location in real-time.
   - Mark stops and delays.
   - Communicate with parents and school admins.

3. **Admin Panel**:
   - Manage buses, drivers, routes, and schedules.
   - Assign children to buses.
   - Monitor overall system performance.

4. **Real-Time Updates**:
   - Location updates via GPS integration.

## Technology Stack
- **Frontend**: HTML, CSS, JavaScript
- **Backend**: Go
- **Database**: SQLite

## Installation
### Prerequisites
- [Go](https://go.dev/) installed
- SQLite installed

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/hakika-ride.git
   cd hakika-ride
   ```

2. Set up the database:
   - Initialize SQLite:
     ```bash
     sqlite3 hakika_ride.db
     ```
   - Use the provided schema to create the tables.

3. Configure the Go backend:
   - Update the `config.json` file with database settings.
   - Install dependencies:
     ```bash
     go mod tidy
     ```

4. Run the backend:
   ```bash
   go run main.go
   ```

5. Frontend setup:
   - Open the `index.html` file in your browser.

6. Test the application:
   - Use the browser-based frontend to interact with the system.

## Usage
1. Register as a parent, driver, or school.
2. Add children and assign them to schools.
3. Monitor buses and receive real-time notifications.

## Future Enhancements
- Integration with GPS hardware for bus location updates.
- Mobile app support (iOS and Android).
- Advanced analytics for school administrators.
- Enhanced security with role-based access control.

## License
This project is licensed under the MIT License.

## Contributors
- [CynthiaOketch](https://github.com/CynthiaOketch)
- [bernotieno](https://github.com/bernotieno)
- [Wambita](https://github.com/Wambita)
- [JosephOkumu](https://github.com/JosephOkumu)

## Acknowledgements
Special thanks to the [SQLite](https://sqlite.org/), [Go](https://go.dev/), and [HTML/CSS/JS] communities for their resources and support.

