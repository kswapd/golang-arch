// static/js/components/tetris.js
const Tetris = {
    props: ['src'],
    template: `
        <div class="external-container">
            <h1>External Page</h1>
            <iframe :src="src" style="width: 100%; height: 500px; border: none;"></iframe>
        </div>
    `
};

window.Tetris = Tetris;