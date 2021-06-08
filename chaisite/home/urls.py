from django.urls import path

from . import views

urlpatterns = [
    path('', views.index, name='index'),
    path('docs/', views.docslist, name="docs-list"),
    path('docs/startup-guide/', views.startup_home, name="startup-home"),
    path('docs/startup-guide/<int:chapter_number>/', views.startup_chapter, name="startup-chapter")
]