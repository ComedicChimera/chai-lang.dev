// mode definition
define('ace/mode/syclone', function(require, exports, module) {
  let oop = require('ace/lib/oop');
  let TextMode = require("ace/mode/text").Mode;
  let SyCloneHighlightRules = require("ace/mode/syclone_highlight_rules").SyCloneHighlightRules;

  let Mode = function() {
    this.HighlightRules = SyCloneHighlightRules;
  }
  oop.inherits(Mode, TextMode);

  exports.Mode = Mode;
});

// rules
define('ace/mode/syclone_highlight_rules', (require, exports, module) => {
  var oop = require("ace/lib/oop");
  var TextHighlightRules = require("ace/mode/text_highlight_rules").TextHighlightRules;

  var SyCloneHighlightRules = function() {

    this.$rules = {
      start: [
        {
          token: 'constant.language',
          regex: /\b(true|false|null)\b/
        },
        {
          token: 'variable.name.syclone',
          regex: /\bthis\b/
        },
        {
          token: 'variable.name.syclone',
          regex: '[\$@][A-Za-z_]\w*'
        },
        {
          token: 'constant.other.byte.syclone',
          regex: /\b(0x[0-9A-F]+|0b[10]+)/
        },
        {
          token: 'constant.numeric.syclone',
          regex: /\b\d+(\.\d+)?i?/
        },
        {
          token: 'string.quoted.double.syclone',
          regex: /\"/,
          next: 'doubleString'
        },
        {
          token: 'comment.line',
          regex: /\/\/[^\n]*/
        },
        {
          token: 'comment.block',
          regex: /\/\*/,
          next: 'multilineComment'
        },
        {
          token: ['period', 'support.function', 'lparen'],
          regex: /(\.)([a-zA-Z_]\w*)\s*(\()/
        },
        {
          token: ['period', 'variable.other'],
          regex: /(\.)([a-zA-Z_]\w*)/
        },
        {
          token: 'variable.other',
          regex: /[@\$][a-zA-Z_]\w*/
        },
        {
          token: 'keyword.operator',
          regex: /[+*\/\-\^=><!&\$@%\:]|(\|\|)/
        },
        {
          token: 'storage.type',
          regex: /\b(list|dict|complex|int|float|byte|bool|str|char)\b/
        },
        {
          token: 'storage.modifier',
          regex: /\b(final|abstract|private|protected|sealed|volatile|global|local|passive|active)\b/
        },
        {
          token: 'keyword.control',
          regex: /\b(if|elif|else|goto|do|for|when|case|select|default|break|continue|return|yield)\b/
        },
        {
          token: 'keyword.other',
          regex: /\b(func|async)\s+/,
          next: 'sycFunction'
        },
        {
          token: 'keyword.other',
          regex: /\b(struct|type|interface)\s+/,
          next: 'dataStructure'
        },
        {
          token: 'keyword.other',
          regex: /\bmodule\s+/,
          next: 'module'
        },
        {
          token: 'keyword.other',
          regex: /\b(use)?include\s+/,
          next: 'include'
        },
        {
          token: 'keyword.other',
          regex: /\b(func|include|use|async|struct|type|interface|await|lambda|value|delete|with|except|throw|constructor|in|new|module)\b/
        },
        {
          token: 'variable.language',
          regex: /this/
        },
        {
          token: ['support.function', 'lparen'],
          regex: /\b(print|println|input|echo|send|listen|decorate|slice|join|split|dllimport|len|range|map|swap|concat|append|reverse|assert|execute)(\()/
        }
      ],
      doubleStrings: [
        {
          token: ''
        }
      ],
      sycFunction: [
        {
          token: 'storage.modifier',
          regex: /(final|abstract|private|protected|sealed|volatile|global|local|passive|active)\s+/
        },
        {
          token: 'entity.name.function',
          regex: /[a-zA-Z_][\w\.]*/,
          next: 'start'
        },
        {
          token: 'lparen',
          regex: /\(/,
          next: 'start'
        }
      ],
      dataStructure: [
        {
          token: 'storage.modifier',
          regex: /(final|abstract|private|protected|sealed|volatile|global|local|passive|active)\s+/
        },
        {
          token: 'entity.name.type',
          regex: /[a-zA-Z_][\w\.]*/,
          next: 'start'
        },
        {
          token: 'lbracket',
          regex: /\{/,
          next: 'start'
        }
      ],
      multilineComment: [
        {
          token: 'comment.block',
          regex: /\*\//,
          next: 'start'
        }, {
          defaultToken: 'comment'
        }
      ],
      module: [
        {
          token: 'storage.modifier',
          regex: /(final|abstract|private|protected|sealed|volatile|global|local|passive|active)\s+/
        },
        {
          token: 'storage.modifier',
          regex: /(passive|active|await)\s+/
        },
        {
          token: 'entity.name.type',
          regex: /[a-zA-Z_][\w\.]*/,
          next: 'start'
        }
      ],
      include: [
        {
          token: 'support.class',
          regex: /[a-zA-Z_][\w\.]*/,
          next: 'start'
        },
        {
          token: 'string.double',
          regex: /"[^\"]*"/,
          next: 'start'
        }
      ]
    }
  }

  oop.inherits(SyCloneHighlightRules, TextHighlightRules);

  exports.SyCloneHighlightRules = SyCloneHighlightRules;
});
