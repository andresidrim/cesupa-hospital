import { Input } from '@/components/Input'
import { FormGroup } from '@/components/FormGroup'

export default function ExamplePage() {
  return (
    <form>
      <FormGroup direction="row" gap="md">
        <Input label="Nome" name="nome" />
        <Input label="CPF" name="cpf" />
      </FormGroup>

      <FormGroup>
        <Input label="Email" name="email" type="email" />
      </FormGroup>
    </form>
  )
}
