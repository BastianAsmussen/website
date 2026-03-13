/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'class',
  content: ['./templates/**/*.html', './static/js/**/*.js'],
  theme: {
    extend: {
      colors: {
        ctp: {
          bg:      'var(--ctp-bg)',
          'bg-alt':'var(--ctp-bg-alt)',
          surface: 'var(--ctp-surface)',
          border:  'var(--ctp-border)',
          text:    'var(--ctp-text)',
          subtext: 'var(--ctp-subtext)',
          muted:   'var(--ctp-muted)',
          accent:  'var(--ctp-accent)',
          teal:    'var(--ctp-teal)',
          green:   'var(--ctp-green)',
          peach:   'var(--ctp-peach)',
          mauve:   'var(--ctp-mauve)',
          red:     'var(--ctp-red)',
          yellow:  'var(--ctp-yellow)',
          sapphire:'var(--ctp-sapphire)',
          sky:     'var(--ctp-sky)',
        },
      },
      fontFamily: {
        mono: ['"JetBrains Mono"', '"Fira Code"', 'Consolas', 'ui-monospace', 'monospace'],
      },
      typography: ({ theme }) => ({
        ctp: {
          css: {
            '--tw-prose-body':        'var(--ctp-text)',
            '--tw-prose-headings':    'var(--ctp-text)',
            '--tw-prose-lead':        'var(--ctp-subtext)',
            '--tw-prose-links':       'var(--ctp-accent)',
            '--tw-prose-bold':        'var(--ctp-text)',
            '--tw-prose-counters':    'var(--ctp-muted)',
            '--tw-prose-bullets':     'var(--ctp-muted)',
            '--tw-prose-hr':          'var(--ctp-border)',
            '--tw-prose-quotes':      'var(--ctp-subtext)',
            '--tw-prose-quote-borders': 'var(--ctp-accent)',
            '--tw-prose-captions':    'var(--ctp-subtext)',
            '--tw-prose-code':        'var(--ctp-teal)',
            '--tw-prose-pre-code':    'var(--ctp-text)',
            '--tw-prose-pre-bg':      'var(--ctp-bg-alt)',
            '--tw-prose-th-borders':  'var(--ctp-border)',
            '--tw-prose-td-borders':  'var(--ctp-border)',
          },
        },
      }),
    },
  },
  plugins: [require('@tailwindcss/typography')],
};
