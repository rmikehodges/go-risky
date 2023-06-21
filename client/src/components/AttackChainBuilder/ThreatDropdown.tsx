import React from 'react';
import { ThreatOutput } from '../Threats/Threats';

interface DropdownProps {
  options: ThreatOutput[] | null;
  onSelectOption: (option: ThreatOutput["name"]) => void;
}

const ThreatDropdown: React.FC<DropdownProps> = ({ options, onSelectOption }) => {

  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedOption = event.target.value;
    onSelectOption(selectedOption);
  }

  return (
    <select onChange={handleSelectChange}>
      {options?.map((option, i) => (
        <option key={i} value={option.name}>{option.name}</option>
      ))}
    </select>
  );
}

export default ThreatDropdown;
