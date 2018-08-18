let demos = ["concur", "func", "generic", "modules", "control-flow"];
let demoNum = 1;

const demoCode = [
    `use include stdio;\nvolatile $counter = 0;\n\nasync printCounter() {\n\tdo (5) Println(counter);\n}\n\nfunc Main() {\n\tprintCounter();\n\tdo (5) counter++;\n}`,
    `func Sum($lst: list[int]) int =>\n\tlst |> ... + |;\n\nfunc Main() {\n\tSum([1, 2, 3]); // 6\n\tSum([4, 5]); // 9\n}`,
    `template&lt;T>\nfunc Max($set: list[T], $fn: func($a: T, $b: T)(bool)) \nlist[T]\n{\n\t$mx: T;\n\tfor set|item| => {\n\t\tif (fn(mx, item))\n\t\t\tmx = item;\n\t}\n}`,
    `\n\nmod Point2D {\n\t$(x, y): int;\n\n\tfunc display() str =>
        "(%d, %d)".format(this.x, this.y);\n}\n\nfunc Main() {\n\t$p = new Point2D { x=2, y=4 };\n\tp.display(); // (2, 4)\n}`,
    `use include stdio;\n\nfunc Main() {\n\t$input = Scanln();\n\tmatch {\n\t\t? len(input) < 5:
        \tPrintln("Input is short.");\n\t\t? input.lower() == "Hello":\n\t\t\tPrintln("Input is a greeting.");\n\t}\n}`
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