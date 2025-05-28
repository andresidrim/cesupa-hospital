'use client'

import React, { useState, useEffect } from 'react'
import './style.css'
import { Input } from '@/components/Input'
import { Button } from '@/components/Button'
import { FormGroup } from '@/components/FormGroup'
import { Select } from '@/components/Select'

interface PacienteFormProps {
  onSubmit: (data: Record<string, string>) => void
  initialData?: Record<string, string>
  isEdit?: boolean
}

export const PacienteForm: React.FC<PacienteFormProps> = ({ onSubmit, initialData, isEdit = false }) => {
  const [form, setForm] = useState({
    nome: '',
    nascimento: '',
    cpf: '',
    sexo: '',
    telefone: '',
    endereco: '',
    email: '',
    tipo_sanguineo: '',
    alergias: '',
  })

  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    if (initialData) {
      setForm({
        nome: initialData.nome || '',
        nascimento: initialData.nascimento || '',
        cpf: initialData.cpf || '',
        sexo: initialData.sexo || '',
        telefone: initialData.telefone || '',
        endereco: initialData.endereco || '',
        email: initialData.email || '',
        tipo_sanguineo: initialData.tipo_sanguineo || '',
        alergias: initialData.alergias || '',
      })
    }
  }, [initialData])

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    const obrigatorios = ['nome', 'nascimento', 'cpf', 'sexo', 'telefone', 'endereco']
    const vazio = obrigatorios.some((campo) => form[campo as keyof typeof form] === '')

    if (vazio) {
      setError('Preencha todos os campos obrigatórios.')
      return
    }

    setError(null)
    onSubmit(form)
  }

  return (
    <form className="paciente-form" onSubmit={handleSubmit}>
      <h2>{isEdit ? 'Editar Paciente' : 'Cadastrar Novo Paciente'}</h2>

      <FormGroup direction="row" gap="md">
        <Input label="Nome completo *" name="nome" value={form.nome} onChange={handleChange} />
        <Input label="Data de nascimento *" name="nascimento" type="date" value={form.nascimento} onChange={handleChange} />
      </FormGroup>

      <FormGroup direction="row" gap="md">
        <Input label="CPF *" name="cpf" value={form.cpf} onChange={handleChange} disabled={isEdit} />
        <Select
          label="Sexo *"
          name="sexo"
          value={form.sexo}
          onChange={handleChange}
          options={[
            { label: 'Masculino', value: 'Masculino' },
            { label: 'Feminino', value: 'Feminino' },
            { label: 'Outro', value: 'Outro' },
          ]}
        />
      </FormGroup>

      <FormGroup direction="row" gap="md">
        <Input label="Telefone *" name="telefone" value={form.telefone} onChange={handleChange} />
        <Input label="Endereço *" name="endereco" value={form.endereco} onChange={handleChange} />
      </FormGroup>

      <FormGroup direction="row" gap="md">
        <Input label="Email" name="email" value={form.email} onChange={handleChange} />
        <Select
          label="Tipo sanguíneo"
          name="tipo_sanguineo"
          value={form.tipo_sanguineo}
          onChange={handleChange}
          options={[
            { label: 'A+', value: 'A+' },
            { label: 'A-', value: 'A-' },
            { label: 'B+', value: 'B+' },
            { label: 'B-', value: 'B-' },
            { label: 'AB+', value: 'AB+' },
            { label: 'AB-', value: 'AB-' },
            { label: 'O+', value: 'O+' },
            { label: 'O-', value: 'O-' },
          ]}
        />
      </FormGroup>

      <FormGroup>
        <Input
    label="Alergias conhecidas"
    name="alergias"
    value={form.alergias}
    onChange={handleChange}
    placeholder="Ex: Lactose, Glúten, Nenhuma"
  />
      </FormGroup>

      {error && <span className="form-error">{error}</span>}

      <div className="button-group">
        <Button type="submit">{isEdit ? 'Salvar Alterações' : 'Cadastrar Paciente'}</Button>
      </div>
    </form>
  )
}
