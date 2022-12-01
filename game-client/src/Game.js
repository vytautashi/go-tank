class Game {
    static GAME_SERVER_URL = "ws://" + document.location.hostname + ":8080/ws";

    static config = {
        type: Phaser.AUTO,
        width: 800,
        height: 600,
        backgroundColor: '#999999',
        parent: 'game-client',
        scale: {
            mode: Phaser.Scale.FIT,
            autoCenter: Phaser.Scale.CENTER_BOTH,
        },
        scene: [SceneGame]
    };

    static game = new Phaser.Game(Game.config);
}

const NOT_FOUND_INDEX = -1;

