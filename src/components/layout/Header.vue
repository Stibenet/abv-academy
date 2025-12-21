<script setup>
import { ref, onBeforeUnmount } from 'vue'

const isMenuOpen = ref(false)

function openMenu() {
  if (!window.matchMedia('(max-width: 800px)').matches) return;
  isMenuOpen.value = true
  document.documentElement.style.overflow = 'hidden'
}


function closeMenu() {
  isMenuOpen.value = false
  document.documentElement.style.overflow = ''
}

function toggleMenu() {
  isMenuOpen.value ? closeMenu() : openMenu()
}

function onKeydown(e) {
  if (e.key === 'Escape') closeMenu()
}

window.addEventListener('keydown', onKeydown)
onBeforeUnmount(() => window.removeEventListener('keydown', onKeydown))
</script>

<template>
  <header class="header">
    <div class="header__container header__container--mobile">
      <!-- Бургер (виден только на мобилке через CSS) -->
      <button class="header__burger" type="button" @click="toggleMenu" aria-label="Открыть меню">
        <span class="header__burger-lines" aria-hidden="true"></span>
      </button>

      <!-- Десктоп-меню (как у тебя было) -->
      <nav class="header__nav header__nav--desktop" aria-label="Главное меню">
        <ul class="header__list">
          <li><a class="header__link header__link--active" href="#">Главная</a></li>
          <li><a class="header__link" href="#">Новости</a></li>
          <li><a class="header__link" href="#">Сведения об образовательной организации</a></li>
          <li><a class="header__link" href="#">Фотогалерея</a></li>
          <li><a class="header__link" href="#">Странички педагогов</a></li>
          <li><a class="header__link" href="#">Контакты</a></li>
        </ul>
      </nav>

      <!-- Кнопка справа: всегда на месте -->
      <button class="header__btn" type="button">
        Личный кабинет
      </button>
    </div>

    <!-- Overlay -->
    <div
        class="mnav-overlay"
        :class="{ 'is-open': isMenuOpen }"
        @click="closeMenu"
        aria-hidden="true"
    />

    <!-- Drawer -->
    <aside class="mnav" :class="{ 'is-open': isMenuOpen }" aria-label="Мобильное меню">
      <div class="mnav__head">
        <div class="mnav__title">Меню</div>
        <button class="mnav__close" type="button" @click="closeMenu" aria-label="Закрыть меню">×</button>
      </div>

      <nav class="mnav__body">
        <a class="mnav__link mnav__link--active" href="#" @click="closeMenu">Главная</a>
        <a class="mnav__link" href="#" @click="closeMenu">Новости</a>
        <a class="mnav__link" href="#" @click="closeMenu">Сведения об образовательной организации</a>
        <a class="mnav__link" href="#" @click="closeMenu">Фотогалерея</a>
        <a class="mnav__link" href="#" @click="closeMenu">Странички педагогов</a>
        <a class="mnav__link" href="#" @click="closeMenu">Контакты</a>
      </nav>
    </aside>
  </header>
</template>
