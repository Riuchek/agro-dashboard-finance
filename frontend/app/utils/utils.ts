function formatValue(value: number, unit: string) {
    if (unit === 'BRL' || unit === 'R$') {
      return new Intl.NumberFormat('pt-BR', {
        style: 'currency',
        currency: 'BRL',
      }).format(value)
    }
  
    return `${new Intl.NumberFormat('pt-BR', {
      minimumFractionDigits: 2,
      maximumFractionDigits: 4,
    }).format(value)} ${unit}`
  }
  
  function formatUpdatedAt(iso: string) {
    if (!iso) {
      return '—'
    }
  
    return new Intl.DateTimeFormat('pt-BR', {
      day: '2-digit',
      month: '2-digit',
      year: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
    }).format(new Date(iso))
  }

export { formatValue, formatUpdatedAt }