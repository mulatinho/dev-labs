
class World { 

    constructor() {
        this.square_w = 48;
        this.canvas = document.getElementById("canvas");
        this.canvas.width = window.innerWidth
        this.canvas.height = window.innerHeight
        this.ctx = canvas.getContext("2d");

        this.ctx.translate(200, 100);
        this.ctx.scale(1, 0.5);
        this.ctx.rotate(Math.PI / 4);

        this.board = [
            ['-','-','-','-','-','-','-','-'],
            ['-','-','-','-','-','-','-','-'],
            ['-','-','-','-','-','-','-','-'],
            ['-','-','-','-','-','-','-','-'],
            ['-','-','-','-','-','-','-','-'],
            ['-','-','-','-','-','-','-','-'],
            ['-','-','-','-','-','-','-','-'],
            ['-','-','-','-','-','-','-','-'],
            ['-','-','-','-','-','-','-','-'],
        ]
        // this.board = [
        //     ['-','-','-','-','-','-','-','-','-','-',],
        //     ['-','-','-','-','-','-','-','-','-','-',],
        //     ['-','-','-','-','-','-','-','-','-','-',],
        //     ['-','-','-','-','-','-','-','-','-','-',],
        //     ['-','-','-','-','-','-','-','-','-','-',],
        //     ['-','-','-','-','-','x','-','-','-','-',],
        //     ['-','-','-','-','-','-','-','-','-','-',],
        //     ['-','-','-','-','-','-','-','-','-','-',],
        //     ['-','-','x','-','-','-','-','-','-','-',],
        //     ['-','-','-','-','-','-','-','-','-','-',],
        //     ['-','-','-','-','-','-','-','-','-','-',],
        // ];

        this.animate()
    }

    animate() {
        var thestr = ""
        this.ctx.clearRect(0,0, this.ctx.innerWidth, this.ctx.innerHeight);
        
        for (var i=0; i<this.board.length; i++) {
            for (var j=0; j<this.board[i].length; j++) {
                if (this.board[i][j] == "-") {
                    this.ctx.fillStyle="blue";
                    this.ctx.fillRect(i*this.square_w, j*this.square_w, this.square_w, this.square_w);
                    this.ctx.strokeRect(i*this.square_w, j*this.square_w, this.square_w, this.square_w);
                    
                    this.ctx.fillStyle="white";
                    thestr = i + "," + j
                    this.ctx.fillText(thestr, (this.square_w*i) + 15, (this.square_w*j) + 25)
                } else {
                    this.ctx.fillStyle="red";
                    this.ctx.fillRect(i*this.square_w, j*this.square_w, this.square_w, this.square_w);
                    this.ctx.strokeRect(i*this.square_w, j*this.square_w, this.square_w, this.square_w);
                    this.ctx.fillStyle="white";
                    thestr = i + "," + j
                    this.ctx.fillText(thestr, (this.square_w*i) + 15, (this.square_w*j) + 25)
                }
            }    
        }
        console.log('animate()')
    }
    
    randomSelection() {
        const i = Math.floor(Math.random() * this.board.length)
        const j = Math.floor(Math.random() * this.board[0].length)
        this.board[i][j] = 'X'
        console.log('randonSelection', i, j)
        console.log(this.board)
        this.animate()
    }
}


function gameApp() {
    window.world = new World();
    window.world.animate()
    // window.world.randomSelection()
}

window.onload = function() { gameApp() }
window.addEventListener('mousedown', function(event) {

    const x = Math.floor(event.x / 65)
    const y = Math.floor(event.y / 65)
    console.log(event)

    console.log(x, y)
    window.world.board[x][y] = 'x'
    window.world.animate()
})