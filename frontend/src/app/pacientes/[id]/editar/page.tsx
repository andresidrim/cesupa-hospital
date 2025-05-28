'use client'

import { useRouter } from 'next/navigation'
import { PageContainer } from '@/components/PageContainer'
import { Header } from '@/components/Header'
import { PacienteForm } from '@/components/PacienteForm'

export default function EditarPacientePage() {
  const router = useRouter()
  const mockPaciente = {
    nome: 'Maria Silva',
    nascimento: '1985-10-12',
    cpf: '123.456.789-00',
    sexo: 'Feminino',
    telefone: '(91) 99999-9999',
    endereco: 'Rua das Flores, 123',
    email: 'maria@exemplo.com',
    tipo_sanguineo: 'O+',
    alergias: 'Nenhuma',
  }

  const handleSubmit = (data: Record<string, string>) => {
    console.log('Paciente atualizado:', data)
    alert('Dados simulados enviados com sucesso!')
    router.push('/pacientes')
  }

  return (
    <PageContainer>
      <Header title="Editar Paciente" actionLabel="Voltar" onActionClick={() => router.push('/pacientes')} />
      <PacienteForm onSubmit={handleSubmit} initialData={mockPaciente} isEdit />
    </PageContainer>
  )
}
