'use client'

import { useParams, useRouter } from 'next/navigation'
import { useEffect, useState } from 'react'
import { PageContainer } from '@/components/PageContainer'
import { Header } from '@/components/Header'
import { Input } from '@/components/Input'
import { Button } from '@/components/Button'
import { Select } from '@/components/Select'
import './style.css'

type Paciente = {
  id: string
  nome: string
  status: string
}

export default function AgendarConsultaPage() {
  const { id } = useParams()
  const router = useRouter()

  const [paciente, setPaciente] = useState<Paciente | null>(null)
  const [form, setForm] = useState({
    data: '',
    horario: '',
    medico: '',
    motivo: '',
  })
  const [erro, setErro] = useState<string | null>(null)

  useEffect(() => {
    const mock: Paciente = {
      id: id as string,
      nome: 'Carlos Andrade',
      status: 'Ativo',
    }
    setPaciente(mock)
  }, [id])

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  const handleSubmit = async () => {
    const { data, horario, medico, motivo } = form
    if (!data || !horario || !medico || !motivo) {
      setErro('Preencha todos os campos obrigatórios.')
      return
    }

    console.log('Consulta agendada (simulado):', {
      pacienteId: id,
      ...form,
    })

    alert('Consulta agendada com sucesso (simulação).')
    router.push('/pacientes')
  }

  if (!paciente) return null

  return (
    <PageContainer>
      <Header
        title="Agendar Consulta"
        actionLabel="Voltar"
        onActionClick={() => router.push(`/pacientes/${id}/editar`)}
      />

      <div className="agendar-card">
        <p><strong>Paciente:</strong> {paciente.nome}</p>
        <p><strong>Status:</strong> {paciente.status}</p>

        <Input
          label="Data da consulta *"
          name="data"
          type="date"
          value={form.data}
          onChange={handleChange}
        />

        <Input
          label="Horário *"
          name="horario"
          type="time"
          value={form.horario}
          onChange={handleChange}
        />

        <Select
          label="Médico responsável *"
          name="medico"
          value={form.medico}
          onChange={handleChange}
          options={[
            { label: 'Dr. Rafael Borges', value: 'rafael_borges' },
            { label: 'Dra. Amanda Leal', value: 'amanda_leal' },
            { label: 'Dr. Luiz Campos', value: 'luiz_campos' },
          ]}
        />

        <Input
          label="Motivo da consulta *"
          name="motivo"
          value={form.motivo}
          onChange={handleChange}
          placeholder="Ex: retorno, avaliação, encaminhamento..."
        />

        {erro && <p className="erro">{erro}</p>}

        <div className="button-group">
          <Button onClick={handleSubmit}>Agendar</Button>
        </div>
      </div>
    </PageContainer>
  )
}
