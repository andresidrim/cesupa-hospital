'use client'

import { useRouter } from 'next/navigation'
import { PageContainer } from '@/components/PageContainer'
import { Header } from '@/components/Header'
import { PacienteConsulta } from '@/components/PacienteConsulta'

export default function PacientesPage() {
  const router = useRouter()

  return (
    <PageContainer>
      <Header
        title="Consultar Paciente"
        actionLabel="Novo Paciente"
        onActionClick={() => router.push('/pacientes/novo')}
      />
      <PacienteConsulta />
    </PageContainer>
  )
}
