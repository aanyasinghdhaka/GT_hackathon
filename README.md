# GT_hackathon

# AI Creative Engine

### *AI-Powered Advertisement Generator | Generative AI & Marketing Tech Track*

Automatically generate **10+ high-quality advertisement creatives** using AI-generated backgrounds, LLM-generated layouts, Fabric.js canvas editing, and full export support.
Upload your **brand logo + product image**, and the system builds **multiple creative variations**, captions, and downloadable assets.

---

# 📌 Table of Contents

1. [Overview](#-overview)
2. [Key Features](#-key-features)
3. [Architecture](#-architecture)
4. [End-to-End Pipeline](#-end-to-end-pipeline)
5. [Tech Stack](#-tech-stack)
6. [Folder Structure](#-folder-structure)
7. [Core Components](#-core-components)
8. [Environment Variables](#-environment-variables)
9. [Running the Project](#-running-the-project)
10. [Why Fabric.js UI?](#-why-fabricjs-ui)
11. [Future Enhancements](#-future-enhancements)

---

# 📌 Overview

Design teams waste **weeks** manually creating ad variants.
Our project solves this with a **fully automated AI Creative Engine** that:

* Accepts a **brand logo + product image**
* Generates **AI-based background variations** (via DALL·E / Gemini Flash Image)
* Produces fabric.js-compatible **layout JSON** for multiple ad sizes
* Renders editable ads in a **canvas editor**
* Generates **captions + ad copy**
* Allows **fine-tuning** and **export**
* Delivers **ZIP downloads** of all creatives

---

# 📌 Key Features

###  1. Auto-Generation of Creative Variants

Product + Logo → AI → multiple ad comps.

### 🎨 2. AI-Generated Backgrounds

Using **DALL·E 3** or **Gemini Flash Image** for background and other artistic elements (IN Process):

* Abstract gradients
* Color palettes
* Themed environments
* Artistic patterns

All stored on **Cloudinary** → fed to layout generator.

### 🧠 3. LLM-Powered Layout Generation

Gemini outputs a **Fabric.js JSON layout** containing:

* Frames
* Shadows
* Decorative blobs
* Gradients
* Typography
* CTA buttons
* Product & logo placement

### 🖼 4. Editable Canvas (Fabric.js)

Users can modify:

* Text
* Font
* Colors
* Images
* Positions
* Sizes

### 📥 5. Export to PNG / ZIP

All creatives are downloadable.

---

# 📌 Architecture

```
User → React UI → Go Backend → Gemini + DALL·E APIs → Cloudinary → DB (optional)
                         ↓
                      Fabric.js JSON
                         ↓
                    Canvas Renderer
                         ↓
              PNG / ZIP Creative Exports
```

---

# 📌 End-to-End Pipeline

## **Step 1 — Upload Assets**

User uploads:

* Brand logo
* Product image

Back-end uploads → **Cloudinary**
Stored and returned as URLs.

---

## **Step 2 — AI Background Generation**

We call:

* DALL·E 3
* OR Gemini Flash Image

Produces **multiple background variations**, each uploaded to Cloudinary.

These URLs are appended to the image list used by the layout model.

DISCLAIMER//NOTE :Our system also supports generating a single, fully-AI-generated advertisement when needed. We already have all required inputs — the product image , brand logo , their LLM-generated descriptions from urls, brand colors, tagline, and creative direction - which are combined into a structured prompt for a full creative ad. This produces a complete image ad via DALL·E 3 or Gemini Image.
However, because full AI image generation can distort logos, alter product appearance, or place incorrect text, we additionally provide a safe mode: AI generates only the background while Fabric.js controls the final layout. This allows users to manually adjust text, CTA, logo placement, and product before the final download, ensuring brand-critical elements remain accurate. Together, this dual-workflow satisfies hackathon requirements (full AI ad generation possible) while still delivering a robust, editable, production-ready creative pipeline.


---

## **Step 3 — Layout Generation (Gemini Flash / Gemini Pro)**

File: `go-api/util/prompts.go`

This file contains:

* `IMAGE_DESCRIPTION_PROMPT`
* `FABRIC_JSON_PROMPT`

Gemini receives:

* Image URLs (logo, product, generated bg)
* Brand colors
* Optional tagline
* Creative direction

Returns **strict JSON** defining:

* instagram_story layout
* instagram_post layout
* facebook_ad layout

Each layout contains `elements[]` (rects, images, text, gradients, shapes).

---

## **Step 4 — Frontend Rendering**

In `Canvas.tsx`:

```ts
canvas.loadFromJSON(layoutJson)
```

Fabric.js renders:

* Layers
* Images
* Text
* Shapes
* Gradients

User can:

* Edit
* Reposition
* Change font
* Change colors
* Modify CTA
* Duplicate creatives

---

## **Step 5 — Exporting**

Frontend supports:

* Export to PNG
* ZIP download
* Per-platform output (Story / Post / Facebook)

---

# 📌 Tech Stack

### **Backend**

* Go 1.22
* Gemini Flash 2.0 / Gemini Pro 1.5
* DALL·E 3 (optional)
* Cloudinary SDK
* Fiber (HTTP)
* PostgreSQL
* Goose migrations
* Docker

### **Frontend**

* React + Vite
* TypeScript
* TailwindCSS
* ShadCN UI
* Zustand
* Fabric.js
* Axios

---

# 📌 Folder Structure

```
canvas-app/
│
├── go-api/
│   ├── main.go
│   ├── handlers/
│   │   ├── apiHandlers.go        # Uploads, generate creative JSON, generate background
│   │
│   ├── util/
│   │   ├── prompts.go            # ⚠ Master prompt file for everything
│   │   ├── cloudinary.go
│   │   ├── dalle.go              # Image generation integration
│   │
│   ├── db/
│       ├── migrations/           # Goose SQL schema
│
│   ├── docker-compose.yml        # Local Postgres
│   ├── .env
│
└── canvas-ui/
    ├── components/
    │   ├── UploadForm.tsx        # Upload assets UI
    │   ├── Canvas.tsx            # Main Fabric.js editor
    │   ├── GeneratedPreview.tsx
    │
    ├── pages/
    │   ├── Generate.tsx
    │
    ├── lib/
        ├── fabricLoader.ts       # Loads Fabric JSON
```

---

# 📌 Core Components

### **1. prompts.go**

This is the brain of the system.

Contains:

* Rules
* Layout schema
* Typography rules
* Gradient rules
* Decorative element rules
* Strict Fabric.js JSON output

This ensures:
✔ Consistent structure
✔ Editable canvas
✔ Multi-platform responsiveness

---

### **2. apiHandlers.go**

Handles:

* Uploads to Cloudinary
* AI image creation
* Prompt construction
* Layout generation via Gemini
* Caption generation
* Response shaping

You can check:

* `GenerateCreativeAdHandler`
* `GenerateImageVariationsHandler`

---

### **3. Canvas.tsx**

Loads Fabric.js JSON:

```ts
canvas.loadFromJSON(layoutJSON)
```

Allows:

* Editing
* Exporting
* Re-coloring
* Moving objects

---

# 📌 Environment Variables

Create `go-api/.env`:

```
GEMINI_API_KEY=
OPENAI_API_KEY=
CLOUDINARY_CLOUD_NAME=
CLOUDINARY_API_KEY=
CLOUDINARY_API_SECRET=
DATABASE_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
```

---

# 📌 Running the Project

### **Backend**

```sh
cd go-api
docker compose up -d
go run main.go
```

### **Frontend**

```sh
cd canvas-ui
npm install
npm run dev
```

Open browser:
👉 [http://localhost:3000](http://localhost:3000)

---

# 📌 Why Fabric.js UI?

Most AI image generators output **flat images**.
Our system outputs **structured, editable design layers**.

This means:

* Marketers can tune spacing
* Change CTA
* Update taglines
* Adapt for campaigns
* Maintain brand identity

This turns the project into a full **Creative Studio**, not just a generator.

---

# 📌 Future Enhancements

✔ Multi-language captions
✔ Inpainting to embed products in scenes
✔ Automatic brand guideline extraction
✔ Auto font detection from brand website
✔ Batch generation (50+ ads)
✔ Model finetuning for ad-specific creativity

---

