import React from 'react';
import { ThreatOutput } from '../Threats/Threats';
import { UUID } from 'crypto';

interface DropdownProps {
  options: ThreatOutput[] | null;
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
        <option key={i} value={option.id}>{option.name}</option>
      ))}
    </select>
  );
}

export default ThreatDropdown;
