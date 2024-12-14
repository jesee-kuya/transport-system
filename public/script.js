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

// Form submissions

// Handle the form submission
document.addEventListener('DOMContentLoaded', function () {
document.querySelectorAll('form').forEach(form => {
  form.addEventListener('submit', async (e) => {
    e.preventDefault();

    const formAction = form.closest('.tab-content').id; // Get login or signup tab
    const formData = new FormData(form); // Extract form data

    console.log(formData)

    const data = {};
    formData.forEach((value, key) => {
      data[key] = value;
    });

    const action = formAction === 'loginContent' ? 'login' : 'signup'; // Decide whether it's signup or login
    try {
      const response = await fetch('/auth', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          action,
          formData: data, // Pass the form data
        }),
      });
      const result = await response.json();
      if (response.ok) {
        alert(result.message);
        // Redirect or show success page
      } else {
        alert(result.error);
      }
    } catch (error) {
      console.error('Error:', error);
      alert('Something went wrong!');
    }
  });
});
});

