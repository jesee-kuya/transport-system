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
