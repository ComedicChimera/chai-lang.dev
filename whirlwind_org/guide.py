import markdown
import os

from .constants import HOST


def load_chapter_bar(chap_num, name):
    chapters = []
    for filename in os.listdir('whirlwind_org/static/markdown/guide'):
        chapter, section_title = filename.split('_')
        chapter = int(chapter)
        if len(chapters) < chapter:
            chapters.append([])
        chapters[chapter - 1].append(section_title[:-3].replace('#', ''))
    chapters[chap_num] = [(x, True) if x == name else x for x in chapters[chap_num]]
    html_elements = []
    for i in range(len(chapters)):
        chapter = chapters[i]
        for j in range(len(chapter)):
            name, selected = chapter[j] if isinstance(chapter[j], tuple) else (chapter[j], False)
            if j == 0:
                element = '<li class="chapter-title%s"><a href="%s/guide/chapter%d"><b>%d</b> %s</a></li>' % (' selected' if selected else '', HOST, i + 1, i + 1, name)
                html_elements.extend([element, []])
            else:
                element = '<li class="chapter-section%s"><a href="%s/guide/chapter%d/%s"><b>%d.%d</b> %s</a></li>' % (
                    ' selected' if selected else '', HOST, i + 1, name, i + 1, j, name
                )
                html_elements[-1].append(element)
    return ''.join(map(lambda x: '<ul>%s</ul>' % ''.join(x) if isinstance(x, list) else x, html_elements))


# flask secures path
def load_guide(chap_num, name):
    with open('whirlwind_org/static/markdown/guide/%d_%s.md' % (chap_num, name)) as file:
        data = file.read()

    html = markdown.markdown(data)
    return html.replace('<code>', '<code class="language-whirlwind">')


def load_chapter_title(chap_num):
    for filename in os.listdir('whirlwind_org/static/markdown/guide'):
        if '#' in filename:
            chapter, _ = filename.split('_')
            if int(chapter) == chap_num:
                with open('whirlwind_org/static/markdown/guide/' + filename) as file:
                    data = file.read()

                html = markdown.markdown(data)
                return html.replace('<code>', '<code class="language-whirlwind">'), filename[:-3].split('#')[1]

