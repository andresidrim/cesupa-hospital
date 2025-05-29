'use client'

import { useEffect, useState } from 'react'
import { useParams, useRouter } from 'next/navigation'
import { PageContainer } from '@/components/PageContainer'
import { Header } from '@/components/Header'
import { Input } from '@/components/Input'
import { Button } from '@/components/Button'
import './style.css'

type PacienteData = {
  id: string
  nome: string
  status: string
  possuiConsultasPendentes: boolean
}

export default function InativarPacientePage() {
  const { id } = useParams()
  const router = useRouter()

  const [paciente, setPaciente] = useState<PacienteData | null>(null)
  const [motivo, setMotivo] = useState('')
  const [erro, setErro] = useState<string | null>(null)

  useEffect(() => {
    // Simulação de carregamento de dados do paciente
    const load = async () => {
      const mock: PacienteData = {
        id: id as string,
        nome: 'Ana Clara',
        status: 'Ativo',
        possuiConsultasPendentes: false, // Troque para true para testar bloqueio
      }
      setPaciente(mock)
    }

    load()
  }, [id])

  const handleSubmit = async () => {
    if (!motivo.trim()) {
      setErro('Informe o motivo da inativação.')
      return
    }

    // Simulação de envio de inativação
    console.log('Paciente inativado:', {
      id,
      motivo,
      data: new Date().toISOString(),
      usuario: 'admin_simulado',
    })

    alert('Paciente inativado com sucesso (simulado).')
    router.push('/pacientes')
  }

  if (!paciente) return null

  return (
    <PageContainer>
      <Header
        title="Inativar Paciente"
        actionLabel="Voltar"
        onActionClick={() => router.push(`/pacientes/${id}/editar`)}
      />

      <div className="inativar-card">
        <p><strong>Nome:</strong> {paciente.nome}</p>
        <p><strong>Status atual:</strong> {paciente.status}</p>

        {paciente.possuiConsultasPendentes ? (
          <p className="alerta">Este paciente possui consultas em aberto. Cancele-as antes de inativar.</p>
        ) : (
          <>
            <Input
              label="Motivo da inativação"
              name="motivo"
              value={motivo}
              onChange={(e) => setMotivo(e.target.value)}
              placeholder="Ex: Transferência, falecimento, duplicidade..."
            />
            {erro && <p className="erro">{erro}</p>}
            <div className="button-group">
              <Button onClick={handleSubmit}>Confirmar Inativação</Button>
            </div>
          </>
        )}
      </div>
    </PageContainer>
  )
}
