import React from 'react';
import  AttackChain  from '../AttackChains/AttackChain';
import { UUID } from 'crypto';

interface DropdownProps {
  options: AttackChain[] | null;
  selectedAttackChain: string;
  onSelectOption: (option: string) => void;
}

const AttackChainDropdown: React.FC<DropdownProps> = ({ options, selectedAttackChain, onSelectOption }) => {

  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedOption = event.target.value;
    onSelectOption(selectedOption);
  }

  return (
    <select value={selectedAttackChain} onChange={handleSelectChange}>
      {options?.map((option, i) => (
        <option key={i} value={option.id?.toString()}>{option.id}</option>
      ))}
    </select>
  );
}

export default AttackChainDropdown;