// Theme toggle: persists in localStorage, defaults to dark.
(function () {
  const html = document.documentElement;
  const stored = localStorage.getItem('theme');
  if (stored === 'light') {
    html.classList.remove('dark');
  } else {
    html.classList.add('dark');
  }
})();

function toggleTheme() {
  const html = document.documentElement;
  const isDark = html.classList.contains('dark');
  if (isDark) {
    html.classList.remove('dark');
    localStorage.setItem('theme', 'light');
  } else {
    html.classList.add('dark');
    localStorage.setItem('theme', 'dark');
  }

  updateToggleIcon();
}

function updateToggleIcon() {
  const isDark = document.documentElement.classList.contains('dark');
  const sunIcon  = document.getElementById('icon-sun');
  const moonIcon = document.getElementById('icon-moon');
  if (!sunIcon || !moonIcon) return;
  if (isDark) {
    sunIcon.classList.add('hidden');
    moonIcon.classList.remove('hidden');
  } else {
    sunIcon.classList.remove('hidden');
    moonIcon.classList.add('hidden');
  }
}

document.addEventListener('DOMContentLoaded', updateToggleIcon);
