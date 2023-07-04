import React from 'react';
import  Threat  from '../Threats/Threat';
import { UUID } from 'crypto';

interface DropdownProps {
  options: Threat[] | null;
  selectedThreat: string;
  onSelectOption: (option: string) => void;
}

const ThreatDropdown: React.FC<DropdownProps> = ({ options, selectedThreat, onSelectOption }) => {

  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedOption = event.target.value;
    onSelectOption(selectedOption);
  }

  return (
    <select value={selectedThreat} onChange={handleSelectChange}>
      {options?.map((option, i) => (
        <option key={i} value={option.id?.toString()}>{option.name}</option>
      ))}
    </select>
  );
}

export default ThreatDropdown;
