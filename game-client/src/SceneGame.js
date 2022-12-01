class SceneGame extends Phaser.Scene {
    conn = null;
    cursors = null;
    players = [];
    bullets = [];
    mapActiveArea = [];

    constructor() {
        super("SceneGame");
    }

    init(data) {
        this.cursors = this.input.keyboard.createCursorKeys();
    }

    preload() {
        this.load.spritesheet('tank', 'assets/tank.png', {
            frameWidth: 64,
            frameHeight: 64
        });
    }

    create() {
        try {
            this.connSetup();
        } catch (error) {
            console.error(error);
        }
    }

    update() {
        if (this.conn.readyState !== WebSocket.OPEN) {
            return;
        }
        this.controls();
    }

    controls() {
        if (this.cursors.left.isDown) {
            this.sendInput(ControlsInputs.MOVE_LEFT);
        }
        else if (this.cursors.right.isDown) {
            this.sendInput(ControlsInputs.MOVE_RIGHT);
        }
        else if (this.cursors.up.isDown) {
            this.sendInput(ControlsInputs.MOVE_UP);
        }
        else if (this.cursors.down.isDown) {
            this.sendInput(ControlsInputs.MOVE_DOWN);
        }

        if (this.cursors.space.isDown) {
            this.sendInput(ControlsInputs.FIRE_GUN);
        }
    }

    connSetup() {
        let self = this;
        this.conn = new WebSocket(Game.GAME_SERVER_URL);
        this.conn.binaryType = 'arraybuffer';

        this.conn.onclose = () => alert("connection lost");

        this.conn.onmessage = function (evt) {
            let b = new Uint8Array(evt.data)
            let commandType = b[0]
            switch (commandType) {
                case CommandFromServer.INIT_PLAYER:
                    MyPlayer.ID = Util.bytesToInt32(b[1], b[2], b[3], b[4]);
                    break;

                case CommandFromServer.INIT_MAP:
                    let mapWidth = Util.bytesToInt16(b[1], b[2]);
                    let mapheight = Util.bytesToInt16(b[3], b[4]);
                    self.initMap(mapWidth, mapheight);

                case CommandFromServer.PLAYERS_POSITIONS:
                    self.updatePositions(
                        Helper.byteMessageToPlayersInfo(b)
                    );
                    break;

                case CommandFromServer.BULLETS_POSITIONS:
                    self.updateBulletsPositions(
                        Helper.byteMessageToBulletsPositions(b)
                    );
                    break;

            }
        };
    }

    initMap(mapWidth, mapheight) {
        // Destroy old map active area graphics
        this.mapActiveArea.forEach(w => w.destroy());
        this.mapActiveArea = [];

        // Adds map active area graphics
        this.mapActiveArea.push(new MapActiveArea(this, mapWidth, mapheight));
    }

    updatePositions(playersInfo) {
        // Updates players positions and creates newly joined players
        playersInfo.forEach(player => {
            const playerIndex = this.players.findIndex(p => p.id == player.id);

            if (playerIndex == NOT_FOUND_INDEX) {
                this.players.push(new Character(this, player.id, player.x, player.y));
            } else {
                this.players[playerIndex].update(player);
            }
        });


        // Destroy left players
        this.players.forEach(player => {
            const playerIndex = playersInfo.findIndex(p => p.id == player.id);
            if (playerIndex == NOT_FOUND_INDEX) {
                let deleteIndex = this.players.findIndex(p => p.id == player.id);
                this.players[deleteIndex].destroy();
                this.players[deleteIndex] = null;
                this.players.splice(deleteIndex, 1);
            }
        });
    }

    updateBulletsPositions(bulletsPositions) {
        // Destroy all bullets
        this.bullets.forEach(bullet => bullet.destroy());
        this.bullets = [];

        // Adds bullets to game
        bulletsPositions.forEach(bullet => {
            this.bullets.push(new Bullet(this, bullet.x, bullet.y));
        });
    }

    // Send information to server about player control inputs (movement, firing gun)
    sendInput(input) {
        let bytearray = new Uint8Array(2);
        bytearray[0] = CommandToServer.UPDATE_PLAYER_INPUT;
        bytearray[1] = input;
        this.conn.send(bytearray.buffer);
    }
}

class MyPlayer {
    static ID = -1;
}

class Character {
    id;
    container;

    constructor(game, id, x = 0, y = 0) {
        this.id = id;

        this.container = game.add.container(x, y, []);
        this.container.add(game.add.sprite(0, 0, 'tank').setName('tank'));


        if (MyPlayer.ID == this.id) {
            game.cameras.main.startFollow(this.container, true, 1.00, 1.00);
        }
    }

    update(player) {
        this.updateAnimation(player);
        this.container.x = player.x;
        this.container.y = player.y;
    }

    updateAnimation(player) {
        if (player.x > this.container.x) {
            this.container.getByName('tank').setFrame(1);
        } else if (player.x < this.container.x) {
            this.container.getByName('tank').setFrame(3);
        } else if (player.y > this.container.y) {
            this.container.getByName('tank').setFrame(2);
        } else if (player.y < this.container.y) {
            this.container.getByName('tank').setFrame(0);
        }
    }

    destroy() {
        this.container.destroy();
    }
}

class Bullet {
    container;

    constructor(game, x = 0, y = 0) {
        this.container = game.add.container(x, y, []);
        this.container.add(game.add.rectangle(0, 0, 4, 4, 0x999999));
    }

    destroy() {
        this.container.destroy();
    }
}

class MapActiveArea {
    container;

    constructor(game, mapWidth, mapheight) {
        this.container = game.add.container(mapWidth / 2, mapheight / 2, []);
        this.container.add(
            game.add.rectangle(0, 0, mapWidth + 54, mapheight + 54, 0x111111)
        );
    }

    destroy() {
        this.container.destroy();
    }
}

// Codes of commands that are sent to server
class CommandToServer {
    static UPDATE_PLAYER_INPUT = 100
}

// Codes of commands that are received from server
class CommandFromServer {
    static INIT_PLAYER = 0;
    static INIT_MAP = 1;
    static PLAYERS_POSITIONS = 2;
    static BULLETS_POSITIONS = 3;
}

class ControlsInputs {
    static MOVE_RIGHT = 0b0000_0001
    static MOVE_LEFT = 0b0000_0010
    static MOVE_DOWN = 0b0000_0100 // y axis inverted in game
    static MOVE_UP = 0b0000_1000 // y axis inverted in game
    static FIRE_GUN = 0b0001_0000
}

class Helper {
    static byteMessageToPlayersInfo(b) {
        let tempPlayersInfo = [];

        for (let i = 1; i < b.length; i += 8) {
            let p = {
                id: Util.bytesToInt32(b[i + 0], b[i + 1], b[i + 2], b[i + 3]),
                x: Util.bytesToInt16(b[i + 4], b[i + 5]),
                y: Util.bytesToInt16(b[i + 6], b[i + 7]),
            };
            tempPlayersInfo.push(p);
        }

        return tempPlayersInfo;
    }

    static byteMessageToBulletsPositions(b) {
        let positions = [];

        for (let i = 1; i < b.length; i += 4) {
            let bullet = {
                x: Util.bytesToInt16(b[i + 0], b[i + 1]),
                y: Util.bytesToInt16(b[i + 2], b[i + 3]),
            };
            positions.push(bullet);
        }

        return positions;
    }
}

class Util {
    static bytesToInt32(b1, b2, b3, b4) {
        return b1 << 24 | b2 << 16 | b3 << 8 | b4
    }

    static bytesToInt16(b1, b2) {
        return b1 << 8 | b2
    }
}

