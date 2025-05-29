'use client'

import React from 'react'
import './style.css'

interface Option {
  label: string
  value: string
}

interface SelectProps extends React.SelectHTMLAttributes<HTMLSelectElement> {
  label?: string
  options: Option[]
  error?: string
}

export const Select: React.FC<SelectProps> = ({ label, options, error, ...props }) => {
  return (
    <div className="select-group">
      {label && <label htmlFor={props.id || props.name}>{label}</label>}
      <select className={`select-field ${error ? 'select-error' : ''}`} {...props}>
        <option value="">Selecione</option>
        {options.map(opt => (
          <option key={opt.value} value={opt.value}>{opt.label}</option>
        ))}
      </select>
      {error && <span className="select-message">{error}</span>}
    </div>
  )
}
