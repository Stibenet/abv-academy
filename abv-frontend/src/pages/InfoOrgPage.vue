<script setup lang="ts">
const file = {
  name: 'DOOP_Erudit_2025-2026_PFDO.pdf',
  url: '/files/DOOP_Erudit_2025-2026_PFDO.pdf'
}
</script>

<template>
  <div class="page">
    <h1>Документы</h1>

    <div class="file-card">
      <div class="file-info">
        <span class="file-name">{{ file.name }}</span>
      </div>

      <a
          :href="file.url"
          target="_blank"
          rel="noopener noreferrer"
          class="file-link"
      >
        Открыть файл
      </a>
    </div>
  </div>
</template>

<style scoped>
.page {
  padding: 24px;
  max-width: 800px;
  margin: 0 auto;
}

.file-card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border: 1px solid #ddd;
  border-radius: 12px;
  padding: 16px;
  margin-top: 16px;
  background: #fff;
}

.file-name {
  font-weight: 500;
}

.file-link {
  background: #2c7be5;
  color: white;
  padding: 8px 14px;
  border-radius: 8px;
  text-decoration: none;
  transition: 0.2s;
}

.file-link:hover {
  background: #1a68d1;
}
</style>