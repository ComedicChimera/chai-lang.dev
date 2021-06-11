from django.shortcuts import render
from django.http import Http404

from .load_content import load_markdown

# Create your views here.
def startup_home(request):
    return render(request, 'docs/guide.html', context={
        'page_content': load_markdown('startup_guide/index'), 
        'page_number': 0
        })

def startup_chapter(request, chapter_number):
    try:
        return render(request, 'docs/guide.html', context={
            'page_content': load_markdown(f'startup_guide/chapter{chapter_number}'), 
            'page_number': chapter_number
            })
    except FileNotFoundError:
        return Http404()