<script setup lang="ts">
const file = {
  name: 'Дополнительная общеобразовательная общеразвивающая программа "Эрудит',
  url: '/files/DOOP_Erudit_2025-2026_PFDO.pdf'
}
</script>

<template>
  <div class="page">
    <h1>Документы</h1>

    <div class="file-card">
      <span class="file-name">{{ file.name }}</span>

      <a
          :href="file.url"
          :download="file.name"
          class="file-link"
      >
        Скачать файл
      </a>
    </div>
  </div>
</template>

<style scoped>
.page {
  padding: 24px;
}

.file-card {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  border: 1px solid #ddd;
  border-radius: 12px;
  padding: 16px;
}

.file-name {
  font-weight: 600;
}

.file-link {
  display: inline-block;
  padding: 10px 14px;
  border-radius: 8px;
  text-decoration: none;
  color: white;
  background: #2c7be5;
}
</style>