from django.urls import path

from . import views

urlpatterns = [
    path('', views.index, name='index'),
    path('docs/', views.docslist, name="docs-list")
]