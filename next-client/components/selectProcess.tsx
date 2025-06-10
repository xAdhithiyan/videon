import { Button, Checkbox, FormControlLabel } from "@mui/material";
import { useState } from "react";

type Items = {
  compress: boolean;
  thumbnails: boolean;
  transcode: boolean;
  watermark: boolean;
  summary: boolean;
};

export default function SelectProcess() {
  const [checkItem, setCheckedItem] = useState<Items>({
    compress: false,
    thumbnails: false,
    transcode: false,
    watermark: false,
    summary: true,
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, checked } = e.target;
    console.log(name, checked);
    setCheckedItem((prev) => ({ ...prev, [name]: checked }));
  };

  const handleSubmit = () => {
    const selectedItems = Object.fromEntries(Object.entries(checkItem).filter(([_, value]) => value));
    console.log(selectedItems);

    // send a fetch request to publish kafka
  };

  return (
    <div className="flex flex-col">
      <FormControlLabel
        control={<Checkbox name="compress" checked={checkItem.compress} onChange={handleChange} />}
        label="Compress Video"
      />
      <FormControlLabel
        control={<Checkbox name="thumbnails" checked={checkItem.thumbnails} onChange={handleChange} />}
        label="Generate Thumbnails"
      />
      <FormControlLabel
        control={<Checkbox name="transcode" checked={checkItem.transcode} onChange={handleChange} />}
        label="TransCode Video"
      />
      <FormControlLabel
        control={<Checkbox name="watermark" checked={checkItem.watermark} onChange={handleChange} />}
        label="Add WaterMark"
      />
      <FormControlLabel
        control={<Checkbox name="summary" checked={checkItem.summary} onChange={handleChange} />}
        label="Video Summary"
      />

      <Button onClick={handleSubmit}>Submit</Button>
    </div>
  );
}
