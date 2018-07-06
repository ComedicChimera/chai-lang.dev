let demos = ["concur", "func", "expr", "modules", "control-flow"];
let demoNum = 1;

const demoCode = [
    `use include stdio;\nvolatile $counter = 0;\n\nasync printCounter() {\n\tdo (5) Println(counter);\n}\n\nfunc Main() {\n\tprintCounter();\n\tdo (5) counter++;\n}`,
    `func Sum($lst: list[int]) int =>\n\tlst |> ... + |;\n\nfunc Main() {\n\tSum([1, 2, 3]); // 6\n\tSum([4, 5]); // 9\n}`,
    `func Sqrt($n: int) {\n\t$t = n;\n\tfor($z = 0; z < 10; z++)\n\t\ty -= (y^2 - n) / (2 * y);\n\treturn y;\n}\n\nfunc Main() {\n\tSqrt(4); // 2\n\tSqrt(16); // 4\n}`,
    `\n\nmodule Point2D {\n\tproperty $(x, y): int;\n\n\tfunc property Display() =>
        "(%d, %d)".Format(this.x, this.y);\n}\n\nfunc Main() {\n\t$p = new Point2D { x=2, y=4 };\n\tp.Display(); // (2, 4)\n}`,
    `use include stdio;\n\nfunc Main() {\n\t$input = Scanln();\n\tmatch {\n\t\t? len(input) < 5:
        \tPrintln("Input is short.");\n\t\t? input.Lower() == "Hello":\n\t\t\tPrintln("Input is a greeting.");\n\t}\n}`
]

$(() => {
    $('#demo-left').mousedown(() => {
        $('#demo-' + demos[0]).removeClass('visible');
        demos.unshift(demos.pop());
        $('#demo-' + demos[0]).addClass('visible');

        if (demoNum > 1) demoNum--;
        else demoNum = 5;
        $('#demo-example-number').html(demoNum);
        
        $('#viewer-content').html(demoCode[demoNum - 1]);
        highlightEditor();
    });

    $('#demo-right').mousedown(() => {
        $('#demo-' + demos[0]).removeClass('visible');
        demos.push(demos.shift());
        $('#demo-' + demos[0]).addClass('visible');

        if (demoNum < 5) demoNum++;
        else demoNum = 1;
        $('#demo-example-number').html(demoNum);
        
        $('#viewer-content').html(demoCode[demoNum - 1]);
        highlightEditor();
    });

    // highlight when page ready
    $('#viewer-content').html(demoCode[0]);
    highlightEditor();
});

function htmlUnescape(string) {
    let doc = new DOMParser().parseFromString(string, "text/html");
    return doc.documentElement.textContent;
}

function highlightEditor() {
    let html = htmlUnescape($('#viewer-content').html());
    $('#viewer-content').html(Prism.highlight(html, Prism.languages.whirlwind));
}