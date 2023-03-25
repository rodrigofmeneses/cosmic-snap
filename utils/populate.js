const url =
  "https://marvelsnapzone.com/getinfo/" +
  "?searchtype=cards&searchcardstype=true";

const trimSpanTag = (text) => {
  const regex = /<\/?span[^>]*>/g;
  return text.replace(regex, "");
};

const getCards = async (url) => {
  try {
    const response = await fetch(url);
    if (!response.ok) {
      console.error("Failed to fetch cards");
      throw new Error(`Failed to fetch cards: ${response.statusText}`);
    }
    const data = await response.json();

    return data.success.cards.map((card) => ({
      name: card.name,
      cost: card.cost,
      power: card.power,
      description: trimSpanTag(card.ability),
      source: card.source,
      image: card.art,
      tags: card.tags.map((tag) => tag.tag),
    }));
  } catch (error) {
    console.error(error);
    throw error;
  }
};

async function main() {
  const cards = await getCards(url);
  for (const card of cards) {
    await fetch("http://localhost:8080/api/v1/card", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(card),
    });
  }
}

main();
