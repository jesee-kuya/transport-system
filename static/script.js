  // Tab switching
  document.querySelectorAll('.tab-btn').forEach(button => {
    button.addEventListener('click', () => {
        document.querySelectorAll('.tab-btn').forEach(btn => btn.classList.remove('active'));
        document.querySelectorAll('.tab-content').forEach(content => content.classList.remove('active'));
        
        button.classList.add('active');
        const tabId = button.dataset.tab + 'Content';
        document.getElementById(tabId).classList.add('active');
    });
});

// Form toggle buttons
document.querySelectorAll('.toggle-btn').forEach(button => {
    button.addEventListener('click', () => {
        const parentDiv = button.closest('.tab-content');
        parentDiv.querySelectorAll('.toggle-btn').forEach(btn => btn.classList.remove('active'));
        parentDiv.querySelectorAll('.form-section').forEach(form => form.classList.remove('active'));
        
        button.classList.add('active');
        const formId = button.dataset.form + 'Form';
        // Find the correct form element
        const formElement = parentDiv.querySelector('#' + formId);
        if (formElement) {
            formElement.classList.add('active');
        }
    });
});
//boarding page scripts
 // Sample student data
 const students = [
    { id: 1, name: 'Alice Brown', grade: '8th', pickup: 'Stop 1', checked: false },
    { id: 2, name: 'James Wilson', grade: '7th', pickup: 'Stop 2', checked: false },
    { id: 3, name: 'Emily Davis', grade: '9th', pickup: 'Stop 3', checked: false },
    { id: 4, name: 'Michael Thompson', grade: '8th', pickup: 'Stop 1', checked: false },
    { id: 5, name: 'Sarah Martinez', grade: '7th', pickup: 'Stop 2', checked: false }
];

// DOM Elements
const studentItems = document.getElementById('studentItems');
const checkedCount = document.getElementById('checkedCount');
const remainingCount = document.getElementById('remainingCount');
const startJourneyBtn = document.getElementById('startJourney');
const mapContainer = document.getElementById('mapContainer');
const notification = document.getElementById('notification');

// Render students
function renderStudents() {
    studentItems.innerHTML = students.map(student => `
        <div class="student-item" data-id="${student.id}">
            <div class="student-info">
                <h4>${student.name}</h4>
                <small>${student.grade} - ${student.pickup}</small>
            </div>
            <div class="student-status">
                <div class="checkbox-wrapper">
                    <div class="custom-checkbox ${student.checked ? 'checked' : ''}" onclick="toggleStudent(${student.id})">
                        <svg viewBox="0 0 24 24" width="16" height="16" fill="currentColor">
                            <path d="M9,16.17L4.83,12l-1.42,1.41L9,19 21,7l-1.41-1.41L9,16.17z"/>
                        </svg>
                    </div>
                </div>
            </div>
        </div>
    `).join('');
}

// Toggle student check-in
function toggleStudent(id) {
    const student = students.find(s => s.id === id);
    if (student) {
        student.checked = !student.checked;
        updateUI();
        showNotification(`${student.name} has ${student.checked ? 'boarded' : 'left'} the bus`);
    }
}

// Update UI
function updateUI() {
    renderStudents();
    const checkedStudents = students.filter(s => s.checked).length;
    checkedCount.textContent = checkedStudents;
    remainingCount.textContent = students.length - checkedStudents;
    startJourneyBtn.disabled = checkedStudents === 0;
}

// Show notification
function showNotification(message) {
    notification.textContent = message;
    notification.classList.add('show');
    setTimeout(() => {
        notification.classList.remove('show');
    }, 3000);
}

// Start journey
startJourneyBtn.addEventListener('click', () => {
    mapContainer.classList.add('active');
    showNotification('Journey started - Parents notified');
    startJourneyBtn.disabled = true;
    startJourneyBtn.textContent = 'Journey in Progress';
    // Here you would normally start the actual tracking
    // and send notifications to parents
});

// Initial render
renderStudents();

// Mobile menu toggle functionality
const menuToggle = document.getElementById('menuToggle');
const sidebar = document.querySelector('.sidebar');

menuToggle.addEventListener('click', () => {
    sidebar.classList.toggle('active');
});

// Close sidebar when clicking outside
document.addEventListener('click', (e) => {
    if (!sidebar.contains(e.target) && !menuToggle.contains(e.target)) {
        sidebar.classList.remove('active');
    }
});

// Close sidebar when window is resized above mobile breakpoint
window.addEventListener('resize', () => {
    if (window.innerWidth > 992) {
        sidebar.classList.remove('active');
    }
});

// Add event listeners for form submission
document.querySelector('form').addEventListener('submit', (e) => {
    e.preventDefault();
    alert('Profile updated successfully!');
});

// Add event listeners for navigation
document.querySelectorAll('.nav-links a').forEach(link => {
    link.addEventListener('click', (e) => {
        e.preventDefault();
        document.querySelectorAll('.nav-links a').forEach(l => l.classList.remove('active'));
        link.classList.add('active');
    });
});

// Mobile menu functionality
const menuBtn = document.querySelector('.menu-btn');
//const sidebar = document.querySelector('.sidebar');

menuBtn.addEventListener('click', () => {
    sidebar.classList.toggle('active');
});

// Close sidebar when clicking outside on mobile
document.addEventListener('click', (e) => {
    if (window.innerWidth <= 768 && 
        !sidebar.contains(e.target) && 
        !menuBtn.contains(e.target) && 
        sidebar.classList.contains('active')) {
        sidebar.classList.remove('active');
    }
});


// Search functionality
const searchInput = document.querySelector('input[type="text"]');
const tableRows = document.querySelectorAll('tbody tr');
//const mapContainer = document.querySelector('.map-container');

// Initialize map section
mapContainer.style.display = 'none';
const mapSection = document.querySelector('.section:nth-child(4)'); 
mapSection.style.display = 'none';

searchInput.addEventListener('input', function(e) {
    const searchTerm = e.target.value.toLowerCase();
    
    tableRows.forEach(row => {
        const text = row.textContent.toLowerCase();
        row.style.display = text.includes(searchTerm) ? '' : 'none';
    });
});

// Navigation active state
document.querySelectorAll('.nav-links a').forEach(link => {
    link.addEventListener('click', (e) => {
        e.preventDefault();
        document.querySelectorAll('.nav-links a').forEach(l => l.classList.remove('active'));
        link.classList.add('active');
    });
});

// Bus selection functionality - Fixed selector
const busTable = document.querySelector('.section:nth-child(3) table'); 
if (busTable) { 
    busTable.querySelectorAll('tbody tr').forEach(row => {
        row.style.cursor = 'pointer';
        row.addEventListener('click', function() {
            // Remove active class from all rows
            busTable.querySelectorAll('tbody tr').forEach(r => {
                r.classList.remove('selected-bus');
            });
            
            // Add active class to clicked row
            this.classList.add('selected-bus');
            
            // Show map section
            mapSection.style.display = 'block';
            mapContainer.style.display = 'flex';
            
            // Scroll to map
            mapSection.scrollIntoView({ behavior: 'smooth' });

            // Get bus details
            const busId = this.querySelector('td:first-child').textContent;
            const route = this.querySelector('td:nth-child(3)').textContent;
            
            // Update map container with bus info
            mapContainer.innerHTML = `
                <div class="map-info">
                    <h3>Live Location</h3>
                    <p>Bus: ${busId}</p>
                    <p>Route: ${route}</p>
                    <div class="map-placeholder">
                        Interactive map will be displayed here showing the location of ${busId}
                    </div>
                </div>
            `;
        });
    });
}
