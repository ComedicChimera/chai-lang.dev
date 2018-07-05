/*
 * SourceLight.js 1.0.0
 * https://github.com/ComedicChimera/SourceLight.js
 * (c) 2018 Jordan Gaines
 * SourceLight may be freely distributed under the MIT License.
 */

function define(name, obj) {
  SourceLight.mappings[name] = obj;
}

function fetch(name) {
  if(!SourceLight.mappings[name]) {
    throw name + " undefined.";
  }
  return SourceLight.mappings[name];
}

class SourceLight {
  static highlight(options) {
    function select(selector) {
      selector = selector || 'code';
      if(selector.startsWith('#')) {
        return [document.getElementById(selector.slice(1))];
      }
      else if(selector.startsWith('.')) {
        return document.getElementsByClassName(selector.slice(1))
      }
      else {
        return document.getElementsByTagName(selector);
      }
    }

    function cssDeserialize(theme) {
      function generateStyles(styleSerial, name, cssMap) {
        let css = "";
        let styleNames = Object.keys(styleSerial);
        let hasStyles = false;
        for(var elem in styleNames) {
          elem = styleNames[elem];
          if(typeof styleSerial[elem] === "string") {
            hasStyles = true;
            css += elem + ":" + styleSerial[elem] + ";";
          }
          else {
            let subMap = {};
            for(var item in elem) {
              generateStyles(styleSerial[elem], elem, subMap);
            }
            for(var subStyle in subMap) {
              cssMap[name + "." + subStyle] = subMap[subStyle];
            }
          }
        }
        if(hasStyles)
          cssMap[name] = css + "\"";
        return cssMap;
      }

      let tokenNames = Object.keys(theme);
      let cssMap = {};
      for(var item in tokenNames) {
        item = tokenNames[item];
        let styleSerial = theme[item];
        cssMap = generateStyles(styleSerial, item, cssMap);
      }
      return cssMap;
    }

    function unescape(text) {
      let unescapeMap = {
          '&amp;': '&',
          '&lt;': '<',
          '&gt;': '>',
          '&quot;': '"',
          '&#x27;': "'",
          '&#x60;': '`'
      };
      for(var special in unescapeMap) {
        var regex = new RegExp(special, 'g');
        if(!regex.test(text))
          continue;
        var component = unescapeMap[special];
        text = text.replace(regex, component);
      }
      return text;
    }

    let elems = select(options.selector);
    options.theme = options.theme || 'default';
    options.mode = options.mode || 'default';
    let cssMap = cssDeserialize(fetch("sourcelight/theme/" + options.theme));
    for(var i = 0; i < elems.length; i++) {
      let item = elems[i];
      item.innerHTML = SourceLight.lex(unescape(item.innerHTML), fetch("sourcelight/mode/" + options.mode), cssMap);
      item.setAttribute("style", cssMap['code.region']);
      item.classList.add('sourcelight-highlight-region');
      item.classList.add('sourcelight-' + options.theme);
    }
  }

  static lex(text, mode, theme) {
    let phrases = {}
    for(var item in mode) {
      let mCase = mode[item];
      let name = mCase.token, regex = mCase.match, sub = mCase.sub;
      if(!regex)
        throw name + "does not have a match case.";
      let matches = text.match(mCase.single ? regex : new RegExp(regex, regex.flags + "g"));
      for(var match in matches) {
        match = matches[match];
        phrases[text.search(regex)] = [match, name, sub];
        text = text.replace(regex, "\0".repeat(match.length), 1);
      }
    }
    let keys = Object.keys(phrases);
    let phraseList = [];
    for(var key in keys) {
      phraseList[key] = phrases[keys[key]];
    }
    for(var phrase in phraseList) {
      phrase = phraseList[phrase];
      text = text.replace("\0".repeat(phrase[0].length), SourceLight.generate(phrase, theme), 1);
    }
    return text;
  }

  static generate(p, theme) {
    if(p[2]) {
      p[0] = SourceLight.lex(p[0], p[2], theme);
    }
    let className = "sourcelight-style-" + p[1].replace(/\./g, "-");
    if(!Object.keys(theme).includes(p[1])) {
      let parents = p[1].split('.');
      while(parents.length > 0) {
        if(Object.keys(theme).includes(parents.join('.'))) {
          return "<span style=\"" + theme[parents.join('.')] + "\" class=\"" + className + "\">" + p[0] + "</span>";
        }
        parents.pop();
      }
      return "<span class=\"" + className + "\">" + p[0] + "</span>";
    }
    return "<span style=\"" + theme[p[1]] + "\" class=\"" + className + "\">" + p[0] + "</span>";
  }
}

SourceLight.mappings = {};

// defaults
define('sourcelight/theme/default', {
  code: {
    region: {
      'border-radius': '5px',
      'background-color': '#fff8f2',
      'box-shadow': 'inset 0 0 2px #212121',
      'display': 'block',
      'font-size': '16px',
      'white-space': 'pre-line',
      'padding': '10px'
    }
  },
  string: {
    single: {
      'color': '#aa0000'
    },
    double: {
      'color': '#aa0000'
    }
  },
  constant: {
    language: {
      'color': '#aa0000'
    },
    numeric: {
      'color': '#0098ba'
    }
  },
  storage: {
    'font-weight': 'bold',
    type: {
      'font-weight': 'bold',
      'color': '#389eff'
    }
  },
  comment: {
    'color': '#666666',
    'font-style': 'italic',
  },
  keyword: {
    'font-weight': 'bold'
  },
  variable: {
    'font-weight': 'bold',
    language: {
      'color': '#B21B39',
      'font-weight': 'bold'
    }
  },
  entity: {
    name: {
      function: {
        'color': '#B21B39',
        'font-weight': 'bold'
      },
      type: {
        'color': '#389eff',
        'font-weight': 'bold'
      }
    }
  }
});

define('sourcelight/mode/default', {});
