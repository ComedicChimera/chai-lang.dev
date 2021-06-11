from django.urls import path

from . import views

urlpatterns = [
    path('guide/', views.startup_home, name="startup-home"),
    path('guide/<int:chapter_number>/', views.startup_chapter, name="startup-chapter")
]
